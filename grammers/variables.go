package main

import "fmt"

func main() {
	structDefinition()
	const tt int = 10
	x := 10
	{
		fmt.Println("value of x:", x)
		x, y := 20, 30
		fmt.Println("value of x:", x)
		fmt.Println("value of y:", y)
	}
	fmt.Println("value of x:", x)
	fmt.Println("value of t:", tt)
	fmt.Println("Result of the calling function multiple:", tt)
	fmt.Println("Result of the calling function simple multiple:", simplemultiple(10, 20))
	a, b := multiple(10, 20)
	fmt.Println("Result of the calling function multiple:", a, b)
	a, b = namedmultiple(10, 20)
	fmt.Println("Result of the calling function named multiple:", a, b)
	a = callByLocaltionMultiple(&x, &x)
	fmt.Println("Result of the calling function callByLocaltionMultiple:", a)

	defer fmt.Println("Deffered function called after the last function being called")
	fmt.Println("This function called before defered function")
	mapDefinition()
	differentWayofSlice()

}
