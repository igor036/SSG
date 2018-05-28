/*
* Author: Igor Joaquim dos  Santos Lima
* E-mail: igorjoaquim.pg@gmail.ccom
*/

package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
)

type Conf struct {
	
	User_Name 	  string
	User_Password string
	Port   		  int
}

func Config() *Conf {

	var conf Conf
	argsCount := 0

	if len(os.Args) == 1 || len(os.Args) < 6 {
		log.Fatal("Arguments were not specified correctly")
		os.Exit(1)
	 }

	 if len(os.Args) > 7 {
		 log.Fatal("Invalid number of arguments")
		 os.Exit(1)
	 }

	 for i := 0; i < len(os.Args); i++ {

		if os.Args[i] == "-usr" {

			i++
			argsCount++
			conf.User_Name = os.Args[i]

		} else if os.Args[i] == "-psw" {
			
			i++
			argsCount++
			conf.User_Password = os.Args[i]
		
		} else if os.Args[i] == "-port" {

			i++
			argsCount++
			port,err := strconv.Atoi(os.Args[i])

			if err != nil || port < 0 {
				log.Fatal("Invalid port number")
				os.Exit(1)
			}

			conf.Port = port
		}
	 }

	 if argsCount < 3 {
		log.Fatal("Some arguments were not specified")
		os.Exit(1)
	 }

	 return &conf
}

func main() {

	if os.Args[1] == "-help" {
		Help()
	}

	conf := Config()
}	

func Help() {
	fmt.Print("-usr -> user name\n-psw -> password\n-port number of port")
	os.Exit(1)
}
