package postfix

import (
	"fmt"
	"strings"
)

// Function to convert postfix to prefix
func PostfixToPrefix(postfix string) (string, error) {
	stack := []string{}
	tokens := strings.Fields(postfix)

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid postfix expression")
			}
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			expression := token + " " + op2 + " " + op1
			stack = append(stack, expression)
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid postfix expression")
	}

	return stack[0], nil
}

func isOperator(token string) bool {
	switch token {
	case "+", "-", "*", "/", "^":
		return true
	}
	return false
}
