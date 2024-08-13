func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	var subset []int
	
	// Sort the array to make sure duplicates are adjacent
	sort.Ints(nums)
	
	backtrack(nums, 0, subset, &result)
	return result
}

// backtrack function generates subsets using backtracking while avoiding duplicates.
func backtrack(nums []int, start int, subset []int, result *[][]int) {
	// Make a copy of the current subset and add it to the result
	temp := make([]int, len(subset))
	copy(temp, subset)
	*result = append(*result, temp)

	for i := start; i < len(nums); i++ {
		// Skip duplicates
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		backtrack(nums, i+1, subset, result)
		subset = subset[:len(subset)-1]
	}
}