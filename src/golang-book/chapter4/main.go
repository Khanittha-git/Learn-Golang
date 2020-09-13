package main

import "fmt"

var catName = "Haji"

func main() {
	var x string
	x = "Hello"
	fmt.Println(x)
	x += ", my name is Tarn"
	fmt.Println(x)
	//อีกวิธี
	//same var y = "Hello World"
	y := "Hello World"
	fmt.Println(y)

	//
	name := "Tarn"
	fmt.Println("My nickname is", name)

	//scope
	//we can define variable outside fucn and it can use in another func
	fmt.Println("My cat's name", catName)

	// constant
	// const is constant variable , can't change the value

	//defining multiple variables
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//input number then *2
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)

}

func cat() {
	fmt.Println("My cat's name", catName)
}
