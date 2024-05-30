package main

func balanceBracket(s string) string {
	var stack []string
	bracketPairMap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	for _, char := range s {
		switch string(char) {
		case "{", "[", "(":
			stack = append(stack, string(char))
		case "}", "]", ")":
			// if found the bracket then pop from the stack
			// stack[len(stack)-1] = the opening bracket from the last stack
			if stack[len(stack)-1] == bracketPairMap[string(char)] {
				stack = stack[:len(stack)-1] // pop from the stack, the last element
			}
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

/**
Complexity
Time complexity: O(n)
Space complexity: O(n)

jumlah looping/operation berdasarkan jumlah karakter dari input string bracket
jumlah penyimpanan `append` stack berdasarkan jumlah karakter yang diinput
*/
