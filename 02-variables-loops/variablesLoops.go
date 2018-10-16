package main

import "fmt"

func main() {
	a, b := 1, 2

	fmt.Printf("a + b %d\n", a + b)

	for a < 100 {
	  a++
	  fmt.Println(a)
	}
}
