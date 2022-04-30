package main

import "fmt"

type employee struct {
	Id          int
	name        string
	family      string
	contactinfo contact
}

type contact struct {
	email   string
	phone   string
	address string
}

func structDefinition() {
	employee1 := employee{1, "Javad", "Bayzavi", contact{"javadbazyavi@gmail.com", "+989173098261", "Jam,Bushehr,IR Iran"}}
	fmt.Println(employee1)        //Print all employee info
	fmt.Println(employee1.family) //access any field of struct

	emp2 := &employee1               //referencing to a struct using pointer
	fmt.Println((*emp2).name, *emp2) //access to a referenced struct member

	emp2.name = "Mohammad"
}
