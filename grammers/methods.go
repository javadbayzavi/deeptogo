package main

import "fmt"

type number int
type degree []string

func (r number) isOdd() bool {
	return r/2 == 0
}

//finding the entry in a slice
func (r degree) indexOf(inputch string) int {
	for index, ele := range r {
		if ele == inputch {
			return index
		}
	}
	return -1
}

//Concatination of two slice using function
func (r degree) conCatinate(indeg degree) degree {
	temp := r
	for x, ele := range indeg {
		temp = append(temp, ele)
	}
	return temp
}

type User struct {
	id         int
	name       string
	age        int
	visitDay   int
	totalLikes int
	followers  int
}

var usr []User = []User{
	User{1, "Javad", 37, 20, 657, 2342},
	User{1, "Ati", 36, 45, 2341, 7000},
	User{1, "Mohammd", 10, 125, 657, 2342},
	User{1, "Ali", 7, 657, 5468, 2342},
}

type By func(user User) bool // Struct type from a function

//Bind a function struct to a method
func (fn By) Filter(us []User) []User {
	filtered := make([]User, 0) //create a slice of struct
	for ele := range us {
		if fn(ele) {
			filtered = append(filtered, ele)
		}
	}
	return filtered
}

type mathCalc func(a int, b int) int //a generic handler for all arithmetic calculation

func (this mathCalc) Compute(a int, b int) int { //Binding a calculation handler to a method
	return this(a, b)
}

func testFuncMethod() {

	var addRes = func(a int, b int) int {
		return a + b
	}
	var minueRes = func(a int, b int) int {
		return a - b
	}
	var MultipleRes = func(a int, b int) int {
		return a * b
	}
	var SubtractRes = func(a int, b int) int {
		return a / b
	}
	var RemindRes = func(a int, b int) int {
		return a % b
	}

	var add = mathCalc(addRes).Compute(6, 3)
	var minues = mathCalc(minueRes).Compute(6, 3)
	var multiple = mathCalc(MultipleRes).Compute(6, 3)
	var subtract = mathCalc(SubtractRes).Compute(6, 3)
	var reminder = mathCalc(RemindRes).Compute(6, 3)

	//custom function 1 and assigned to a variable
	var ageaccepted = func(user User) bool {
		if user.age >= 20 {
			return true
		} else {
			return false
		}
	}
	//custom function 1 and assigned to a variable
	var followersaccepted = func(user User) bool {
		if user.followers >= 1500 {
			return true
		} else {
			return false
		}
	}

	var handler1 = By(ageaccepted).Filter(usr)

	fmt.Print(handler1)
}
