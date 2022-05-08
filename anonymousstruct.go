package main

func structcreator(name string) struct {
	id    int
	age   float32
	speed float32
} {
	//return _{id: 1, age:36, speed:120}
	type ss struct {
		id    int
		age   float32
		speed float32
	}
	return test(ss{id: 1, age: 38.4, speed: 120})
}

func test(v struct {
	id    int
	age   float32
	speed float32
}) struct {
	id    int
	age   float32
	speed float32
} {
	return v
}
