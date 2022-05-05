package main

import (
	"fmt"
)

//Generic Builder interface for general purpose builder
type BuilderInterface interface {
	BuildPart(partname string)
	GetProduct() ProductInterface
	SetProduct(value ProductInterface)
	Reset()
}

//Implementation of a Product builder
type Builder struct {
	product ProductInterface
}

func (this *Builder) Reset() {
	this.product = &Product{}
}
func (this *Builder) GetProduct() ProductInterface {
	return this.product
}
func (this *Builder) SetProduct(value ProductInterface) {
	this.product = value
}
func (this *Builder) BuildPart(partname string) {
	this.GetProduct().AddPart(partname)
}

//Generic Interface for a configurable product
type ProductInterface interface {
	GetParts() []string
	AddPart(partname string)
}

//Implementation of a Product that a builder can build
type Product struct {
	parts []string
}

func (this *Product) GetParts() []string {
	return this.parts
}
func (this *Product) AddPart(partname string) {
	this.parts = append(this.parts, partname)
}

//Builder planning engine act as a Abstract Factory to receive a blueprint and build products
type BuildPlanner struct {
	builder Builder
}

func (this BuildPlanner) BuildWithPlan(plan string) ProductInterface {
	//Initiate a new Builder engine
	this.builder = Builder{}

	//Reset the product line to default
	this.builder.Reset()

	if plan == "A" {
		this.builder.BuildPart("Door")
		this.builder.BuildPart("Brake")
		this.builder.BuildPart("Wheels")
	} else {
		this.builder.BuildPart("Door")
	}
	return this.builder.GetProduct()
}

func testBuilder() {
	var productFactory = BuildPlanner{}
	var product = productFactory.BuildWithPlan("A")

	fmt.Print("Blueprint for Product A: ")
	fmt.Println(product.GetParts())

	product = productFactory.BuildWithPlan("B")
	fmt.Print("Blueprint for Product B: ")
	fmt.Println(product.GetParts())

}
