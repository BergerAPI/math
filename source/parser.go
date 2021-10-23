package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// All tokens
var tokens []TokenMatch

// The current token index
var currentIndex = -1

// The current token
var currentToken TokenMatch

// advancing the current token
func nextToken() (TokenMatch, error) {
	if currentIndex+1 >= len(tokens) {
		return TokenMatch{}, errors.New("the index is bigger than the length of the tokens")
	}

	currentIndex++

	return tokens[currentIndex], nil
}

// Expecting a certain token
func expect(tokenType int) TokenMatch {
	if currentToken.tokenType != tokenType {
		fmt.Println("Expected token with type of " + string(rune(tokenType)) + ", got " + string(rune(currentToken.tokenType)))
		os.Exit(2)
	}

	token, err := nextToken()

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	currentToken = token

	return currentToken
}

// Math factor: Int, Float, LParen, RParen
func factor() (Node, error) {
	switch currentToken.tokenType {
	case IntegerLiteral:
		asInt, err := strconv.Atoi(currentToken.raw)

		// Somehow, the int isn't an int
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		expect(IntegerLiteral)

		return IntegerLiteralNode{value: asInt}, nil

	case Identifier:
		expect(Identifier)

		return VariableAccessNode{name: currentToken.raw}, nil
	}

	return nil, errors.New("didn't find this token")
}

// Math term
func term() Node {
	node, _ := factor()

	for currentToken.tokenType != EndOfLine && (currentToken.raw == "*" || currentToken.raw == "/") {
		operator := currentToken.raw

		expect(ArithmeticOperator)

		right, _ := factor()

		node = ExpressionNode{operator: operator, left: node, right: right}
	}

	return node
}

// A basic math equation
func expression() Node {
	expr := term()

	for currentToken.tokenType != EndOfLine && (currentToken.raw == "+" || currentToken.raw == "-") {
		operator := currentToken.raw

		expect(ArithmeticOperator)

		expr = ExpressionNode{operator: operator, left: expr, right: term()}
	}

	return expr
}

// Starting to parse the tokens into an AST
func parse(_tokens []TokenMatch) Node {
	tokens = _tokens

	currentToken, _ = nextToken()
	expr := expression()

	return expr
}
