package main

import "fmt"

//Generice interface for all command context
type CommandInterface interface {
	Run(context CommandInterface, name string)
	processMe(params []string)
}

//Abstract parent class that bundle common behavior of all command context
type CommandParent struct {
}

//common behavior of all command context
func (this CommandParent) Run(context CommandInterface, name string) {
	context.processMe(append(make([]string, 0), name))
}

//Concerete ChildCommand1 command context with set of specific behaviors
type ChildCommand1 struct {
	CommandParent
	classname string
}

//Gateway to specific behaviors through command
func (this *ChildCommand1) processMe(params []string) {
	this.classname = params[0]
	//specific strategy for this child class
	fmt.Println(this.showMyInfo())
}

//Specific behavior of a child command context
func (this *ChildCommand1) showMyInfo() string {
	return this.classname
}

type ChildCommand2 struct {
	CommandParent
}

//Gateway to specific behaviors through command
func (this *ChildCommand2) processMe(params []string) {
	//specific strategy for this child class
	fmt.Println(this.extraSpecific())
}

//Specific behavior of a child command context
func (this *ChildCommand2) extraSpecific() string {
	return "This is extract specific behavior for ChildCommand2"
}

func testCommand() {
	engine1 := ChildCommand1{}
	engine1.Run(&engine1, "ChildCommand1 class")
	engine2 := ChildCommand2{}
	engine2.Run(&engine2, "")
}
