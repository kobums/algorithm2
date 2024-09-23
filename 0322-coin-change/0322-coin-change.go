func coinChange(coins []int, amount int) int {
	// dp 배열을 amount+1 크기로 초기화하고, 모든 값을 큰 값으로 설정 (amount+1)
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	// 기저 조건: 0원을 만들기 위해 필요한 동전의 개수는 0
	dp[0] = 0

	// 각 금액에 대해 최소 동전 개수 계산
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// dp[amount]가 여전히 초기값이면, 금액을 만들 수 없는 경우이므로 -1 반환
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// Helper function to return the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}