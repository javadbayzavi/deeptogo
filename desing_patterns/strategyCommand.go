package main

import "fmt"

//Generice interface for all command context
type parentinterface interface {
	Run(context parentinterface, name string)
	processMe(params []string)
}

//Generic interface for all executor engine
type executorInterface interface {
	execute(params []string)
}
type executor struct {
	context parentinterface
}

func (this executor) execute(params []string) {

	this.context.processMe(params)
}

//Abstract parent class that bundle common behavior of all command context
type parent struct {
	executor executorInterface
}

//common behavior of all command context
func (this parent) Run(context parentinterface, name string) {
	this.executor = executor{context: context}
	this.executor.execute(append(make([]string, 0), name))
}

//Concerete childType1 command context with set of specific behaviors
type childType1 struct {
	parent
	classname string
}

//Gateway to specific behaviors through command
func (this childType1) processMe(params []string) {
	this.classname = params[0]
	//specific strategy for this child class
	fmt.Println(this.showMyInfo())
}

//Specific behavior of a child command context
func (this childType1) showMyInfo() string {
	return this.classname
}

type childType2 struct {
	parent
}

//Gateway to specific behaviors through command
func (this childType2) processMe(params []string) {
	//specific strategy for this child class
	fmt.Println(this.extraSpecific())
}

//Specific behavior of a child command context
func (this childType2) extraSpecific() string {
	return "This is extract specific behavior for childType 2"
}

func testStrategyCommand() {
	engine1 := childType1{}
	engine1.Run(engine1, "child class")
	engine2 := childType2{}
	engine2.Run(engine2, "")
}
