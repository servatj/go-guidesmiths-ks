package main

import "fmt"

func square(i int) int {
  return i*i
}

func main() {
  for i := 0; i < 100; i++ {
  	fmt.Println(i, square(i))
  }
}
