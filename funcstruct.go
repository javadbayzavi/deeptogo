package main

import "fmt"

type netPrice func(price float64, ratePay float64) float64

func factorCalculator(itemprices []float64, discont float64, fn netPrice) {

	for index, item := range itemprices {
		fmt.Printf("Net Price for Item %d is: %d", index, fn(item, discont))
	}
}
func cacula() {
	var f1 netPrice = func(p float64, r float64) float64 {
		return p * r / 100
	}

	var items []float64 = []float64{2500, 3200, 1540, 20000, 12540}
	factorCalculator(items, 0.10, f1)
}
