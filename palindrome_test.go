package main

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	got := isPalindrome("katak")

	if got != true {
		t.Errorf("Test case failed. Expected %v, got %v", true, got)
	}
	fmt.Println(got)
}
