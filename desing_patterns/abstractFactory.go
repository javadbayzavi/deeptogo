package main

import (
	_ "fmt"
)

type AbstratFactoryInterface interface {
	create_conceret_entity(value string)
}

type AbstractFactory struct {
}

func (this *AbstractFactory) create_conceret_entity(value string) entity_interface {
	if value == "A" {
		return &conceret_entityA{}
	} else {
		return &conceret_entityB{}
	}
}

type entity_interface interface {
	showMyInfo() string
}

type conceret_entityA struct {
}

func (this *conceret_entityA) showMyInfo() string {
	return "This message from conceret entity A"
}

type conceret_entityB struct {
}

func (this *conceret_entityB) showMyInfo() string {
	return "This message from conceret entity B"
}

func createEntity() {
	var factory = AbstractFactory{}
	var entity = factory.create_conceret_entity("A")
	println(entity.showMyInfo())
	entity = factory.create_conceret_entity("B")
	println(entity.showMyInfo())
}

func main() {
	createEntity()
}
