package main

func highestPalindrome(s string, k int) string {
	n := len(s)
	changes := make([]bool, n)

	// Convert string to slice of bytes for mutability
	bytes := []byte(s)

	// make it a palindrome
	if !makePalindrome(bytes, changes, 0, n-1, &k) {
		return "-1"
	}

	// maximize the palindrome value
	upgradePalindrome(bytes, changes, 0, n-1, &k)

	return string(bytes)
}

func makePalindrome(bytes []byte, changes []bool, left, right int, k *int) bool {
	if left >= right {
		return true
	}

	if bytes[left] != bytes[right] {
		if *k <= 0 {
			return false
		}
		// Choose the maximum of the two to make both ends equal
		if bytes[left] > bytes[right] {
			bytes[right] = bytes[left]
		} else {
			bytes[left] = bytes[right]
		}
		changes[left] = true
		changes[right] = true
		*k--
	}

	return makePalindrome(bytes, changes, left+1, right-1, k)
}

func upgradePalindrome(bytes []byte, changes []bool, left, right int, k *int) {
	if left >= right || *k <= 0 {
		return
	}

	if bytes[left] == bytes[right] {
		if bytes[left] != '9' {
			if changes[left] || changes[right] {
				if *k >= 1 {
					bytes[left] = '9'
					bytes[right] = '9'
					*k--
				}
			} else {
				if *k >= 2 {
					bytes[left] = '9'
					bytes[right] = '9'
					*k -= 2
				}
			}
		}
	} else {
		if bytes[left] != '9' && bytes[right] != '9' && *k >= 1 {
			bytes[left] = '9'
			bytes[right] = '9'
			*k--
		}
	}

	upgradePalindrome(bytes, changes, left+1, right-1, k)
}
