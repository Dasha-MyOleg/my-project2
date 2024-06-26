package main

import (
	"my-project2/postfix"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
	tests := []struct {
		postfix    string
		expected   string
		shouldFail bool
	}{
		{"4 2 - 3 * 5 +", "+ * - 4 2 3 5", false},
		{"2 3 + 4 *", "* + 2 3 4", false},
		{"5 6 2 + * 12 4 / -", "- * 5 + 6 2 / 12 4", false},
		{"", "", true},
		{"2 3 + +", "", true},
	}

	for _, test := range tests {
		result, err := postfix.PostfixToPrefix(test.postfix)
		if test.shouldFail {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}
