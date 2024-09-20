func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// Case 1: Exclude the first house and rob from house 2 to the last house
	case1 := robLinear(nums[1:])

	// Case 2: Exclude the last house and rob from house 1 to the second-to-last house
	case2 := robLinear(nums[:len(nums)-1])

	return max(case1, case2)
}

// robLinear solves the "House Robber" problem for a linear arrangement of houses
func robLinear(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}

	return dp[n-1]
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}