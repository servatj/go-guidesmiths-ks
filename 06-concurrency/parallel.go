package main

import (
	"fmt"
	"time"
)

// sleeps for `secs` seconds
func f1(secs time.Duration) (result string) {
	fmt.Printf("waiting %v\n", secs)
	time.Sleep(secs * time.Second)
	result = fmt.Sprintf("waited for %v seconds", secs)
	return
}

// prints arg1, arg2
func f2(arg1, arg2 string) {
	fmt.Println(arg1)
	fmt.Println(arg2)
}

// this function executes for 3 seconds, because waits a lot
func runNotParallel() {
	out1 := f1(2)
	out2 := f1(1)
	f2(out1, out2)

}

// golang parallel return functions
// todo: make it run so all the function will executes for 2 seconds not for 3
func runParallel() {
	out1 := make(chan string)
	out2 := make(chan string)
	go func() {
		out1 <- f1(2)
	}()
	go func() {
		out2 <- f1(1)
	}()
	f2(<-out1, <-out2)
}

func main() {
	runNotParallel()
	runParallel()
}
