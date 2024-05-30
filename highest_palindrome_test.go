package main

import (
	"testing"
)

func TestHighestPalindrome(t *testing.T) {

	testCases := []struct {
		palindrome string
		k          int
		expected   string
	}{
		{
			palindrome: "radar",
			k:          1,
			expected:   "radar",
		},
		{
			palindrome: "9899",
			k:          2,
			expected:   "9999",
		},
		{
			palindrome: "abcd",
			k:          1,
			expected:   "-1",
		},
		{
			palindrome: "932239",
			k:          2,
			expected:   "992299",
		},
		{
			palindrome: "1234567",
			k:          2,
			expected:   "-1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.palindrome, func(t *testing.T) {
			result := highestPalindrome(tc.palindrome, tc.k)
			if result != tc.expected {
				t.Errorf("Test case %s failed. Expected %v, got %v", tc.palindrome, tc.expected, result)
			}
		})
	}
}
