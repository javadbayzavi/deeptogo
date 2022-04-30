package main

func multiple(x int, y int) (int, int) {
	return x * y, ((y + x) / 100)
}

func simplemultiple(x int, y int) int {
	return x * y
}

func namedmultiple(x int, y int) (add int, complex int) {
	add = x + y
	complex = add * x * y
	return
}

//A form of call by reference. Goland does not support call by reference however we can use the concept of pointers
func callByLocaltionMultiple(x *int, y *int) int {
	return (*x) * (*y)
}

//Variable numer of function parameters
func variadicmultiple(x int, y ...int) int {
	return x * y[len(y)-1]
}
