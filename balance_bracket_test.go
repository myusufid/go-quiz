package main

import (
	"testing"
)

func Test_balanceBracket(t *testing.T) {

	testCases := []struct {
		brackets string
		expected string
	}{
		{
			brackets: "{ ( ( [ ] ) [ ] ) [ ] }",
			expected: "YES",
		},
		{
			brackets: "{ [ ( ] ) }",
			expected: "NO",
		},
		{
			brackets: "{ [ ( ) ] }",
			expected: "YES",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.brackets, func(t *testing.T) {
			result := balanceBracket(tc.brackets)
			if result != tc.expected {
				t.Errorf("Test case %s failed. Expected %v, got %v", tc.brackets, tc.expected, result)
			}
		})
	}
}
