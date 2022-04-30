package main

import (
	"fmt"
	"sort"
)

func differentWayofSlice() {
	var slice1 = []int{1, 2, 3, 4} //slice def using var

	slice2 := []int{4, 5, 6, 7} //short hand slice def
	println("slice 1: ", slice1)
	println("slice 2: ", slice2)

	var array = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"} //full array for slice creation
	var slcie3 []string = array[4:4]                                  //cut a slice from array
	var slcie4 []string = slcie3[2:3]                                 //create a slice frm another slice
	var slice5 = make([]int, 5)                                       //create a slice using make command

	sort.Strings(slcie3) //sort slice using sort library 'imported'
	sort.Ints(slice1)    //again sort slice using sort library 'imported'
	fmt.Println(slcie4)
	fmt.Println(slice5)

}
