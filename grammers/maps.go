package main

import (
	"fmt"
)

func mapDefinition() {
	var firstmap map[string]int
	fmt.Println("First map is:", firstmap)

	var secondmap map[string]int = map[string]int{"A": 10, "B": 20, "C": 30, "D": 40, "E": 50, "F": 60, "G": 70}
	fmt.Println("Second map is:", secondmap)

	var thirdmap = map[string]int{"A": 10, "B": 20, "C": 30, "D": 40, "E": 50, "F": 60, "G": 70}
	fmt.Println("Third map is:", thirdmap)

	var forthmap = make(map[string]int)
	fmt.Println("Forth map is:", forthmap)

	for key, value := range secondmap {
		println("Key and value :", key, value)
	}

	forthmap["AA"] = 10 //Adding a pair to map

	dd := forthmap["AA"] //Retreiveing the value of a pair in map
	fmt.Print(dd)

	key, value := forthmap["AA"] //check the existence of a pair in map
	fmt.Println(value, key)

	delete(forthmap, "AA")      //delete a pair from map
	newrefrencedmap := forthmap //maps are reference type, and change in one pair will affect on other maps
	fmt.Println(newrefrencedmap)

}
