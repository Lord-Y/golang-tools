package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		value    string
		array    []string
		expected bool
	}{
		{
			value:    "value",
			array:    []string{"a", "b", "value"},
			expected: true,
		},
		{
			value:    "v",
			array:    []string{"a", "b", "value"},
			expected: false,
		},
	}

	for _, tc := range tests {
		z := StringInSlice(tc.value, tc.array)
		assert.Equal(tc.expected, z)
	}
}

func TestSliceDifference(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		value1   []string
		value2   []string
		expected []string
	}{
		{
			value1:   []string{"foo", "bar", "hello"},
			value2:   []string{"foo", "bar"},
			expected: []string{"hello"},
		},
		{
			value1:   []string{"foo", "bar", "hello"},
			value2:   []string{},
			expected: []string{"foo", "bar", "hello"},
		},
	}

	for _, tc := range tests {
		z := SliceDifference(tc.value1, tc.value2)
		assert.Equal(tc.expected, z)
	}
}
