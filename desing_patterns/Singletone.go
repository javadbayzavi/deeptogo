package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type SingletoneInterface interface {
	createInstance() Instance
}

type Singletone struct {
	instanse *Instance
}

func (this *Singletone) createInstance() Instance {
	if this.instanse == nil {
		lock.Lock()
		defer lock.Unlock()
		this.instanse = &Instance{Instancename: "Single Instance"}
	}
	return *this.instanse
}

type Instance struct {
	Instancename string
}

func (this *Instance) print() {
	fmt.Println("Instance name is:" + this.Instancename)
}

func testSingletone() {
	var singlecreator = Singletone{}
	var instance1 = singlecreator.createInstance()
	fmt.Print("Info of instance1:")
	instance1.print()

	var instance2 = singlecreator.createInstance()
	fmt.Print("Info of instance2:")
	instance2.print()

}
