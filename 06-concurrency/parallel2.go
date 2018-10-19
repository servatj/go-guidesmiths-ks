package main

import "sync"

// Parallelize parallelizes the function calls
func Parallelize(functions ...func()) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))

	defer waitGroup.Wait()

	for _, function := range functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}
}

func main() {
	value1 := ""
	value2 := ""

	func1 := func() {
		value1 = f1(2)
	}

	func2 = func() {
		value2 = f1(1)
	}

	Parallelize(func1, func2)

	f2(out1, out2)
}
