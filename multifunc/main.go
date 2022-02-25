// Package multifunc illustrates how Go uses goroutines and channels to make
// concurrency a first class citizen.
//
// In this second iteration, the code is leverages a WaitGroup to provide
// synchronization between two goroutines; resulting in a non-blocking code
// where the two function calls run concurrently.
//
// IMPORTANT: When working with concurrent code, involved goroutines must NOT
// share mutable state, as it may cause inconsistencies / bugs. Also, output
// won't be deterministic (meaning the output may be out of order sometimes)
// as each function may run out of order achording to the underlying OS scheduler.
//
// Sample Output:
//  Waiting for goroutines to finish...
//  Generating number 1
//  Generating number 2
//  Printing number 1
//  Generating number 3
//  Printing number 2
//  Printing number 3
//  Done!
//
// Source: https://www.digitalocean.com/community/tutorials/how-to-run-multiple-functions-concurrently-in-go
package main

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, wg *sync.WaitGroup) {
	// Decrease WaitGroup count by one after function ends.
	defer wg.Done()
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
	}
}

func printNumbers(wg *sync.WaitGroup) {
	// Decrease WaitGroup count by one after function ends.
	defer wg.Done()
	for idx := 1; idx <= 3; idx++ {
		fmt.Printf("Printing number %d\n", idx)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	// We expect Done to be called twice as we are calling two functions.
	wg.Add(2)

	// Creating a goroutine is as simple as prepending the calling function with
	// the `go` keyword.
	go printNumbers(wg)
	go generateNumbers(3, wg)

	fmt.Println("Waiting for goroutines to finish...")
	// If we remove the Wait() call, code may not fully execute at all, as the
	// main goroutine may finish before the other subroutines end theirs.
	wg.Wait()
	fmt.Println("Done!")
}
