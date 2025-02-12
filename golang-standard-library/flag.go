// Package flag berisikan fungsionalitas untuk memparsing command line argument
// https://golang.org/pkg/flag/
package main

import (
	"flag"
	"fmt"
)

func main() {
	 username  := flag.String("username", "root","database username")
	 password  := flag.String("password", "root","database password")
	 host  := flag.String("host", "localhost","database host")
	 port  := flag.Int("port", 0,"database port")

	//  var username *string = flag.String("username", "root","database username")
	//  var password *string  = flag.String("password", "root","database password")
	//  var host *string  = flag.String("host", "localhost","database host")
	//  var port *int  = flag.Int("port", 0,"database port")

	 flag.Parse()

	 fmt.Println("Username",*username)
	 fmt.Println("Password",*password)
	 fmt.Println("Hosting",*host)
	 fmt.Println("Port",*port)




}

