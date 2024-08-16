func partition(s string) [][]string {
	var result [][]string
	var currentPartition []string

	// Helper function for backtracking
	var backtrack func(start int)
	backtrack = func(start int) {
		if start >= len(s) {
			// Add a copy of currentPartition to the result
			temp := make([]string, len(currentPartition))
			copy(temp, currentPartition)
			result = append(result, temp)
			return
		}

		for end := start; end < len(s); end++ {
			if isPalindrome(s, start, end) {
				// If the substring is a palindrome, add it to the currentPartition
				currentPartition = append(currentPartition, s[start:end+1])
				// Recurse with the next substring
				backtrack(end + 1)
				// Backtrack by removing the last element
				currentPartition = currentPartition[:len(currentPartition)-1]
			}
		}
	}

	// Start the backtracking process
	backtrack(0)

	return result
}

// isPalindrome checks if a substring of s from start to end is a palindrome.
func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}