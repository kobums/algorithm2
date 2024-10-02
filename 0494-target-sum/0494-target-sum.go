func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    
    // target이 sum 범위 안에 있지 않으면 불가능
    if target > sum || target < -sum {
        return 0
    }
    
    s := (target + sum) / 2
    
    if (target+sum) % 2 != 0 || s < 0 {
        return 0
    }
    
    dp := make([]int, s + 1)
    dp[0] = 1
    
    for _, num := range nums {
        for i := s; i >= num; i-- {
            dp[i] += dp[i - num]
        }
    }
    
    return dp[s]
}