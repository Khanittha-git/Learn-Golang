package main

import "fmt"

func main() {
	//Ex.1
	var x [5]int
	x[4] = 100 // defined 5nd in array x is 100
	fmt.Println(x)
	// result be [0 0 0 100]
	fmt.Println(x[4])
	// result be 100

	//Ex.2
	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64 = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	//fmt.Println(total / len(y))
	// should chang to same type because total is float64 but len(y) is int
	fmt.Println(total / float64(len(y)))
	// result be 86.6

	//another way to solve
	var total2 float64 = 0
	// use _ for tell to compiler is the _ variable is not used
	for _, value := range y {
		total2 += value
	}
	fmt.Println(total / float64(len(y)))

	//create array
	m := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	fmt.Println(m)

	//2.Slice

}
