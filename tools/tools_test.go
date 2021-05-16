package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		value          string
		array          []string
		expected       bool
		failureMessage string
	}{
		{
			value:          "value",
			array:          []string{"a", "b", "value"},
			expected:       true,
			failureMessage: "Fail to get expected boolean",
		},
		{
			value:          "v",
			array:          []string{"a", "b", "value"},
			expected:       false,
			failureMessage: "Fail to get expected boolean",
		},
	}

	for _, tc := range tests {
		z := StringInSlice(tc.value, tc.array)
		assert.Equal(tc.expected, z)
	}
}
