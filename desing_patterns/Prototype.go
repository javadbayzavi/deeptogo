package main

import "fmt"

type PrototypeInterface interface {
	clone() PrototypeInterface
	print()
}

type Prototype struct {
	property1 string
	property2 int
	classname string
}

func (this *Prototype) clone() PrototypeInterface {
	var clonedobj = &Prototype{
		property1: this.property1,
		property2: this.property2,
		classname: this.classname + "_cloned",
	}
	return clonedobj
}

func (this *Prototype) print() {
	fmt.Println("Classname :" + this.classname + " Property1:" + this.property1 + " Property2: " + string(this.property2))
}

func testPrototype() {
	prototyper := Prototype{
		classname: "Prototype",
		property1: "Value",
		property2: 20,
	}
	var prototype = prototyper.clone()

	prototyper.print()
	prototype.print()
}
