func change(amount int, coins []int) int {
    dp := make([]int, amount+1)
    dp[0] = 1 // 0원을 만드는 방법은 아무 동전도 사용하지 않는 것 1가지

    for _, coin := range coins {
        for i := coin; i <= amount; i++ {
            dp[i] += dp[i-coin]
        }
    }
    
    return dp[amount]
}