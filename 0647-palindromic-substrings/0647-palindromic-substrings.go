func expandAroundCenter(s string, left, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}

// Function to count the number of palindromic substrings
func countSubstrings(s string) int {
	totalPalindromes := 0

	for i := 0; i < len(s); i++ {
		// Count odd-length palindromes (centered at s[i])
		totalPalindromes += expandAroundCenter(s, i, i)

		// Count even-length palindromes (centered between s[i] and s[i+1])
		totalPalindromes += expandAroundCenter(s, i, i+1)
	}

	return totalPalindromes
}