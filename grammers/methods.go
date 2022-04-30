package main

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
		temp[len(temp)+x] = ele
	}
	return temp
}
