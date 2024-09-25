func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// dp 배열을 1로 초기화 (최소한 각 숫자는 자기 자신만의 부분 수열을 가짐)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	// 각 숫자에 대해 이전 숫자들과 비교하여 증가하는 부분 수열을 찾음
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// dp 배열에서 가장 큰 값이 최장 증가 부분 수열의 길이
	maxLength := 0
	for _, length := range dp {
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}