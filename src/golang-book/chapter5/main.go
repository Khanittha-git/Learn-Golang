package main

import "fmt"

func main() {
	//print 1-10
	//for
	for i := 1; i <= 10; i++ {
		fmt.Println(i)

	}

	//if
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
	//switch
	t := 2
	fmt.Println("Write", t, "as")
	switch t {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	for x := 1; x <= 100; x++ {
		if x%3 == 0 {
			fmt.Println(x)
		}
	}

	//FizzBuzz
	/*divisible by 3 retustn Fizz , divisible by 7 retustn FizzBuzz
	divisible by 3 and 5 return FizzBuzz*/
	for x := 1; x <= 100; x++ {
		if (x%3 == 0) && (x%5 == 0) {
			fmt.Println("FizzBuzz")
		} else if x%5 == 0 {
			fmt.Println("Buzz")
		} else if x%3 == 0 {
			fmt.Println("Fizz")
		}
	}

}
