func leastInterval(tasks []byte, n int) int {
	// Frequency map to count occurrences of each task
	taskFreq := make([]int, 26)
	for _, task := range tasks {
		taskFreq[task-'A']++
	}

	// Sort task frequencies in descending order
	sort.Slice(taskFreq, func(i, j int) bool {
		return taskFreq[i] > taskFreq[j]
	})

	// Maximum frequency of any task
	maxFreq := taskFreq[0]
	// Count how many tasks have this maximum frequency
	maxCount := 0
	for _, freq := range taskFreq {
		if freq == maxFreq {
			maxCount++
		}
	}

	// The formula to calculate the minimum intervals needed
	partCount := maxFreq - 1
	partLength := n - (maxCount - 1)
	emptySlots := partCount * partLength
	availableTasks := len(tasks) - maxFreq*maxCount
	idles := max(0, emptySlots-availableTasks)

	return len(tasks) + idles
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}