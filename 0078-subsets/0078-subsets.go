func subsets(nums []int) [][]int {
    result := [][]int{}
    subset := []int{}
    backtrack(nums, 0, subset, &result)
    return result
}

// backtrack function uses backtracking to generate all possible subsets.
func backtrack(nums []int, start int, subset []int, result *[][]int) {
    // Make a copy of the current subset and add it to the result.
    temp := make([]int, len(subset))
    copy(temp, subset)
    *result = append(*result, temp)
    
    for i := start; i < len(nums); i++ {
        subset = append(subset, nums[i])
        backtrack(nums, i+1, subset, result)
        subset = subset[:len(subset)-1] // backtrack
    }
}