func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // DP 배열 생성 및 1로 초기화
    dp := make([]int, len(nums))
    for i := range dp {
        dp[i] = 1
    }

    maxLen := 1

    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            // 현재 숫자가 이전 숫자보다 크다면
            if nums[i] > nums[j] {
                if dp[j]+1 > dp[i] {
                    dp[i] = dp[j] + 1
                }
            }
        }
        if dp[i] > maxLen {
            maxLen = dp[i]
        }
    }

    return maxLen
}