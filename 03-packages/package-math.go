package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(math.Pow(float64(i),2))
	}
}
