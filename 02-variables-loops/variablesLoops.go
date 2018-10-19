package main

import "fmt"

func main() {
	a, b := 1, 2

	fmt.Printf("a + b %d\n", a+b)

	fmt.Printf("a is of type %T\n", a)

	for a < 100 {
		a++
		fmt.Println(a)
	}

	var x int = 5
	var y int = 10
	var sum int = x + y

	fmt.Println(sum)

}
