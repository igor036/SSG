/*
* Author: Igor Joaquim dos  Santos Lima
* E-mail: igorjoaquim.pg@gmail.ccom
*/

package main

import (
  "fmt"
  "os"
  "log"
)

//color's
const RED   string = "\033[31m"
const GREEN string = "\033[32m"
const BLUE  string = "\033[34m"
const NAME  string = "Ssg@Client:"

func PrintErr(msg string,err error, kill bool) {

  if (err != nil) {
    fmt.Printf(BLUE+"%s %s %s Error: %s\n",NAME,RED,msg,err);
  } else {
    fmt.Printf(BLUE+"%s %s %s\n",NAME,RED,msg);
  }
  
  if (kill) {
    os.Exit(1)
  }
}

func Print(msg string) {
  fmt.Printf("%s %s %s %s\n",BLUE,NAME,GREEN,msg)
}

func Logger(msg string) {
  log.Printf("%s %s\n",BLUE,msg)
}