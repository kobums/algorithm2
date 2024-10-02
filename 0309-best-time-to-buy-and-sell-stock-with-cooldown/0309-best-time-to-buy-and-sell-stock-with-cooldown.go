func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    n := len(prices)
    hold := make([]int, n)
    sold := make([]int, n)
    cooldown := make([]int, n)
    
    hold[0] = -prices[0] // 첫날 주식을 산 경우
    sold[0] = 0          // 첫날 주식을 팔 수 없으므로 0
    cooldown[0] = 0      // 첫날 쿨다운도 불가능하므로 0
    
    for i := 1; i < n; i++ {
        hold[i] = max(hold[i-1], cooldown[i-1] - prices[i])
        sold[i] = hold[i-1] + prices[i]
        cooldown[i] = max(cooldown[i-1], sold[i-1])
    }
    
    // 마지막 날에 주식을 팔거나 쿨다운 상태였을 때의 최대 이익
    return max(sold[n-1], cooldown[n-1])
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}