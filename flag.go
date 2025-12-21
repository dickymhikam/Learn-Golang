package main

import (
	"flag"
	"fmt"
)

// Package Flag adalah berisikan fungsionalitas untuk memparsing command line argument

func main()  {
	 var username *string = flag.String("username", "root","database username")
	 var password *string = flag.String("password", "root","database password")
	 var host *string = flag.String("host", "localhost","database host")
	 var port *int = flag.Int("port", 0,"database port")

	flag.Parse()

	 fmt.Println("username", *username)
	 fmt.Println("password", *password)
	 fmt.Println("host", *host)
	 fmt.Println("port", *port)
	}