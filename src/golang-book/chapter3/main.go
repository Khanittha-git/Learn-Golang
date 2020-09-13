package main

import "fmt"

func main() {
	//number
	fmt.Println("int 1+1 = ", 1+1)
	fmt.Println("float 1+1 = ", 1.0+1.0)
	//string
	fmt.Println(len("Hello World")) //11
	fmt.Println("Hello World"[1])   //101 [1] result is 2nd character
	fmt.Println("Hello " + "World")
	//booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
