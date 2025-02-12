package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element array",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Multiple elements array",
			input:    []int{38, 27, 43, 3, 9, 82, 10},
			expected: []int{3, 9, 10, 27, 38, 43, 82},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := mergeSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("mergeSort(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
