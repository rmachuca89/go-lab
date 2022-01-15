package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	// Defining a boolean flag `-w` to count words
	words = flag.Bool("w", true, "Count words")

	// Defining a boolean flag `-l` to count lines
	lines = flag.Bool("l", false, "Count lines")

	// Defining a boolean flag `-c` to count bytes
	btes = flag.Bool("c", false, "Count bytes")
)

func count(r io.Reader, countLines bool, countWords bool, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files or STDIN)
	scanner := bufio.NewScanner(r)

	// Set count type to scan. Mutually exclusive as they override each other,
	// having bytes with the highest priority.
	if countLines {
		scanner.Split(bufio.ScanLines)
	}

	if countWords {
		scanner.Split(bufio.ScanWords)
	}

	// Note that a string is a sequence of bytes, not runes.
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	//Defining an internal scoped word counter
	counter := 0

	// For every word scanned, increment the counter by one
	for scanner.Scan() {
		counter++
	}

	// Return the total counter
	return counter
}

func main() {
	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words (or lines)
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines, *words, *btes))
}
