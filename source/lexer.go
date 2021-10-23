package main

import (
	"strconv"
)

// Basically all tokens, with their identifier
const (
	IntegerLiteral     = 1
	ArithmeticOperator = 2
	Identifier         = 3
)

// TokenMatch There is a match.
type TokenMatch struct {
	tokenType int
	raw       string
}

// IsNumeric Checking is a string is numeric.
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// isArithmeticOperator Checking if we have an "+", "-", "*", "/"
func isArithmeticOperator(char string) bool {
	return char == "+" || char == "-" || char == "*" || char == "/"
}

// lex Lexing a string, and matching tokens.
func lex(line string) []TokenMatch {
	var results []TokenMatch
	index := 0

	for index < len(line) {
		char := string([]rune(line)[index])

		// We have a number. That means, increase until the number ends.
		if IsNumeric(char) {
			raw := char

			for index+1 < len(line) {
				currentChar := string([]rune(line)[index+1])

				if IsNumeric(currentChar) {
					index++
					raw += currentChar
				} else {
					break
				}
			}

			results = append(results, TokenMatch{tokenType: IntegerLiteral, raw: raw})
		} else if isArithmeticOperator(char) {
			results = append(results, TokenMatch{tokenType: ArithmeticOperator, raw: char})
		} else {

			// Identifier
			if !isArithmeticOperator(char) && !IsNumeric(char) && char != " " {
				raw := char

				for index+1 < len(line) {
					currentChar := string([]rune(line)[index+1])

					if isArithmeticOperator(currentChar) || IsNumeric(currentChar) || currentChar == " " {
						break
					}

					index++
					raw += currentChar
				}

				index++
				results = append(results, TokenMatch{tokenType: Identifier, raw: raw})
			}
		}

		index++
	}

	return results
}
