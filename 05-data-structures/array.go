package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var books [4]string

	books[0] = "Book 0"
	books[1] = "Book 1"
	books[2] = "Book 2"
	books[3] = "Book 3"

	fmt.Println(books)

	numbers := []int{1, 2, 3, 4}

	for index, value := range numbers {
		fmt.Println("index ", index, "value ", value)
	}

}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
