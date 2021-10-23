package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * Reading a line from an reader
 */
func readline(r *bufio.Reader) (string, error) {
	var (
		isPrefix       = true
		err      error = nil
		line, ln []byte
	)

	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return string(ln), err
}

/**
 * Basic main function
 */
func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome To math. Please input your math equation.")

	input(in)
}

/**
 * Scanning what the user typed in and then lexing, parsing and interpreting.
 */
func input(reader *bufio.Reader) {
	line, _ := readline(reader)
	results := lex(line)

	for _, value := range results {
		fmt.Println(value)
	}

	// Calling the function again.
	input(reader)
}
