// Package multifunc illustrates how Go uses goroutines and channels to make
// concurrency a first class citizen.
//
// In this first iteration, the code is executed secuentially and blocking;
// if hypotetically each call took 3 seconds to run, it would take 6 seconds
// total for main to complete its execution.
//
// Sample Output:
//  Waiting for goroutines to finish...
//  Printing number 1
//  Printing number 2
//  Printing number 3
//  Generating number 1
//  Generating number 2
//  Generating number 3
//
// Source: https://www.digitalocean.com/community/tutorials/how-to-run-multiple-functions-concurrently-in-go
package main

import (
	"fmt"
)

func generateNumbers(total int) {
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
	}
}

func printNumbers() {
	for idx := 1; idx <= 3; idx++ {
		fmt.Printf("Printing number %d\n", idx)
	}
}

func main() {
	printNumbers()
	generateNumbers(3)
}
