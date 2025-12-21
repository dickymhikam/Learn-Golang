package main

import "fmt"

func main()  {
	firstname := "Dicky"
	lastname := "M"

	fmt.Println("Hello '", firstname, lastname ,"'")
	fmt.Printf("Hello '%s %s'\n", firstname, lastname)
}