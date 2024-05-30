package main

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		// input[0] = a
		// input[4-0-1] = 2
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
