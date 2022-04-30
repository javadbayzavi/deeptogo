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

type By func(user User) bool

func (fn By) Filter(us []User) []User {
	filtered := make([]User, 0)
	for ele := range us {
		if fn(ele) {
			filtered = append(filtered, ele)
		}
	}
}
