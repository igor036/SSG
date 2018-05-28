/*
* Author: Igor Joaquim dos  Santos Lima
* E-mail: igorjoaquim.pg@gmail.ccom
*/

package main

import (
	
	"os"
	"log"
	"net"
	"errors"
	"strconv"
	"golang.org/x/crypto/ssh"
)

type SSH_Server struct {

	User_Name     string
	User_Password string
	Host_Address  string
	Host_Port     int
}

type Server_connection struct {

	Connection *ssh.Client
	Session    *ssh.Session
}

func ServerConfig() (*SSH_Server,error) {

	var server SSH_Server

	if len(os.Args) == 1 || len(os.Args) <= 8 {
		return nil,errors.New("Arguments were not specified correctly")
	}
	
	argsCount := 0
	for i := 0; i < len(os.Args); i++ {	

		if os.Args[i] == "-usr" {

			i++
			argsCount++
			server.User_Name = os.Args[i]

		} else if os.Args[i] == "-psw" {

			i++
			argsCount++
			server.User_Password = os.Args[i]

		} else if os.Args[i] == "-addr" {

			i++
			argsCount++
			server.Host_Address = os.Args[i]

		} else if os.Args[i] == "-port" {

			i++
			argsCount++
			port,err := strconv.Atoi(os.Args[i])
			if (err != nil) {
				return nil,errors.New("Port value is not valid!")
			}
			server.Host_Port = port
		}
	}

	if argsCount < 4 {
		return nil, errors.New("Arguments invalid!")
	}

	return &server,nil
}

func SshConfig(usr string, psw string) *ssh.ClientConfig {

	if usr == "" || psw == "" {
		return nil
	}

	return  &ssh.ClientConfig{
		User: usr,
		Auth: []ssh.AuthMethod {
			ssh.Password(psw),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
}

func Connect(server *SSH_Server, config *ssh.ClientConfig) (*Server_connection,error)  {

	strAddr := server.Host_Address + ":" + strconv.Itoa(server.Host_Port)
	connection,err := ssh.Dial("tcp",strAddr,config)

	if (err != nil) {
		return nil,err
	}
	defer connection.Close()

	session,err := connection.NewSession()
	if (err != nil || Xterm(session) != nil) {
		return nil,err
	}
	defer session.Close()

	return &Server_connection {
		Connection: connection,
		Session: session,
	},nil
}

func Xterm(session *ssh.Session) error {

	modes := ssh.TerminalModes {
		ssh.ECHO:          0,    
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		session.Close()
		return err
	}

	session.Stdout = os.Stdout
    session.Stderr = os.Stderr
    session.Stdin = os.Stdin
    session.Shell()
    session.Wait()

	return nil
}

func main() {

	server,err := ServerConfig()
	if err != nil {
		log.Fatal("Error on config server! %s\n",err)
	}

	config := SshConfig(server.User_Name,server.User_Password)
	if config == nil {
		log.Fatal("Error in create sshClientConfig %s\n",err)
	}

	_,err = Connect(server,config)
	if err != nil {
		log.Fatal("Error on connection %s\n",err)
	}
}
