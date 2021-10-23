package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
)

/**
 * Reading a line from a reader
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
 * Returning the value of something.
 */
func solveNode(node Node) (int, error) {
	switch reflect.TypeOf(node).Name() {
	case "ExpressionNode":
		expr, ok := node.(ExpressionNode)

		if !ok {
			fmt.Println("Something went miserably wrong.")
			os.Exit(2)
		}

		left, _ := solveNode(expr.left)
		right, _ := solveNode(expr.right)

		switch expr.operator {
		case "+":
			return left + right, nil
		case "-":
			return left - right, nil
		case "*":
			return left * right, nil
		case "/":
			return left / right, nil
		}
	case "IntegerLiteralNode":
		literal, ok := node.(IntegerLiteralNode)

		if !ok {
			fmt.Println("Something went miserably wrong.")
			os.Exit(2)
		}

		return literal.value, nil
	}

	return 0, errors.New("there is no node like this")
}

/**
 * Scanning what the user typed in and then lexing, parsing and interpreting.
 */
func input(reader *bufio.Reader) {
	fmt.Print("Your equation: ")

	line, _ := readline(reader)
	result := parse(lex(line))

	expr, ok := result.(ExpressionNode)

	// If the result isn't an expression, we can ignore it.
	if ok {
		solved, err := solveNode(expr)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(solved)
	} else {
		fmt.Println("Nothing to solve.")
	}

	// Calling the function again.
	input(reader)
}
