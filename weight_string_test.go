package main

import (
	"reflect"
	"testing"
)

func TestWeightedStrings(t *testing.T) {

	testCases := []struct {
		word     string
		queries  []int
		expected []string
	}{
		{
			word:     "abbcccd",
			queries:  []int{1, 3, 9, 8},
			expected: []string{"Yes", "Yes", "Yes", "No"},
		},
		{
			word:     "gits",
			queries:  []int{7, 9, 20, 19},
			expected: []string{"Yes", "Yes", "Yes", "Yes"},
		},
		{
			word:     "abc",
			queries:  []int{1, 2, 3},
			expected: []string{"Yes", "Yes", "Yes"},
		},
		{
			word:     "xyz",
			queries:  []int{24, 25, 26},
			expected: []string{"Yes", "Yes", "Yes"},
		},
		{
			word:     "abcabc",
			queries:  []int{1, 3, 6, 7},
			expected: []string{"Yes", "Yes", "No", "No"},
		},
		{
			word:     "abc",
			queries:  []int{4, 5, 6},
			expected: []string{"No", "No", "No"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.word, func(t *testing.T) {
			result := WeightedStrings(tc.word, tc.queries)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Test case %s failed. Expected %v, got %v", tc.word, tc.expected, result)
			}
		})
	}
}
