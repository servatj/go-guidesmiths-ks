package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	books := make(map[int]string)

	books[1] = "book 1"
	books[2] = "book 365"

	for key, value := range books {
		fmt.Println("key ", key, "value ", value)
	}

}
