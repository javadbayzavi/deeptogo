package main

import "fmt"

func ifcondition() {
	if 10 == 10 {
		fmt.Println("Call for If condtion")
	} else if 10 < 11 {
		fmt.Println("Call for else if condtion")
	} else {
		fmt.Println("Call for else condtion")
	}

}

func switchFunc(day int) {
	//day := 8;
	switch day {
	case 1:
		fmt.Println("case 1 in switch has issued")
	case 2:
		fmt.Println("case 2 in switch has issued")
	case 3:
		fmt.Println("case 1 in switch has issued")
	case 4:
		fmt.Println("case 2 in switch has issued")
	case 5:
		fmt.Println("case 1 in switch has issued")
		fallthrough
	case 6:
		fmt.Println("case 2 in switch has issued both with fallthrough and it 6 case")
	default:
		fmt.Println("default case in switch has issued")
	}
}

func onlyForLoop() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println("The calculation using for loop for sum 1 - 10 is =", sum)
labelname:
	for i := 0; i <= 10; i++ {
		if i == 6 {
			break
		} else if i == 7 {
			goto labelname
		} else if i == 8 {
			continue
		}
		fmt.Println("Loop for checking the break and continue cluse")
	}

}
