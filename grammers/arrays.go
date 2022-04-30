package main

import "fmt"

func arraydefinition() {
	//array definition using var keyword
	var firstarray []int
	fmt.Println(firstarray)

	//Shorhand definition of arrays
	secondarray := []int{1, 2, 3, 4, 5, 6, 78, 9}
	fmt.Println(secondarray)
	for i := 0; i < len(secondarray); i++ {
		fmt.Println(secondarray[i])
	}

	var partialarray [5]int = [5]int{22, 44}  //first partially init of array
	partialarray = [5]int{33, 22, 33, 44, 55} //re-init of the array with full size data
	fmt.Println(partialarray)

	b := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9} // variadic array at init
	var c [4]int = [4]int{0: 12, 1: 22, 2: 33, 3: 44}                   // init of array by its indices
	c[2] = 13                                                           // set array element value

	for value := range b {
		fmt.Println(value) // a simple form of loop for array without index
	}
	passArrya(secondarray) //Notice: we cannot send partialarray, because of its predefined fixed length

}
func passArrya(arr []int) {
	//Do some thing with passed array
}
