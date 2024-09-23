func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1 // return valid bounds
}

// Function to return the longest palindromic substring
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// Check for odd-length palindromes (center is s[i])
		left1, right1 := expandAroundCenter(s, i, i)
		// Check for even-length palindromes (center is between s[i] and s[i+1])
		left2, right2 := expandAroundCenter(s, i, i+1)

		// Update the start and end pointers if a longer palindrome is found
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}

	return s[start : end+1]
}