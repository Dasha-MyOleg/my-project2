package main

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/Dasha-MyOleg/my-project2/postfix"
	"github.com/stretchr/testify/assert"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (h *ComputeHandler) Compute() error {
	inputBytes, err := io.ReadAll(h.Input)
	if err != nil {
		return err
	}

	postfixExpr := strings.TrimSpace(string(inputBytes))
	if postfixExpr == "" {
		return fmt.Errorf("input expression is empty")
	}

	result, err := postfix.PostfixToPrefix(postfixExpr)
	if err != nil {
		return err
	}

	_, err = io.WriteString(h.Output, result)
	return err
}

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		input      string
		expected   string
		shouldFail bool
	}{
		{"4 2 - 3 * 5 +", "+ * - 4 2 3 5\n", false},
		{"2 3 + 4 *", "* + 2 3 4\n", false},
		{"", "", true},
		{"2 3 + +", "", true},
	}

	for _, test := range tests {
		input := strings.NewReader(test.input)
		output := &strings.Builder{}
		handler := &ComputeHandler{
			Input:  input,
			Output: output,
		}

		err := handler.Compute()
		if test.shouldFail {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, strings.TrimSpace(output.String()))
		}
	}
}
