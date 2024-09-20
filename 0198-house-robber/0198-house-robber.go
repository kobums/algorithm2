func rob(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// dp[i] will represent the maximum amount of money that can be robbed up to house i
	dp := make([]int, len(nums))

	// Base cases
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	// Fill the dp array
	for i := 2; i < len(nums); i++ {
		// You can either rob this house and add it to dp[i-2], or skip it and keep dp[i-1]
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}

	// The last element in dp array contains the maximum amount of money that can be robbed
	return dp[len(nums)-1]
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}