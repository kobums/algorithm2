func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	// dp[i] will store the minimum cost to reach step i
	dp := make([]int, n+1)

	// Base cases: The cost to start at step 0 or step 1 is 0, since you don't pay until you move forward
	dp[0] = 0
	dp[1] = 0

	// Fill the dp array
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	// The minimum cost to reach the top is either from the last or second to last step
	return dp[n]
}

// Helper function to get the minimum of two numbers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}