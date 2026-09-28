package main

import "fmt"

func main (){
	var name string

	name = "Yahya al wahab"
	fmt.Println("Nama : ", name)

	var lastname string
	
	lastname = "wahab"
	fmt.Println("lastname:", name, lastname)
	
	middlename := "al"
	fmt.Println("nama tengah:", middlename)

	var (
		fullName = "Yahya al wahab"
		firstName = "Yahya"
	)

	fmt.Println(fullName)
	fmt.Println(firstName)

	
}