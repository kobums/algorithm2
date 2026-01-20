func maxCoins(nums []int) int {
    n := len(nums)
    
    // 양 끝에 1 추가
    arr := make([]int, n+2)
    arr[0], arr[n+1] = 1, 1
    for i, v := range nums {
        arr[i+1] = v
    }
    
    // dp[i][j] = i~j 범위에서 최대 코인
    dp := make([][]int, n+2)
    for i := range dp {
        dp[i] = make([]int, n+2)
    }
    
    // 구간 길이를 늘려가며 계산
    for length := 1; length <= n; length++ {
        for left := 1; left <= n-length+1; left++ {
            right := left + length - 1
            
            for k := left; k <= right; k++ {
                coins := arr[left-1] * arr[k] * arr[right+1]
                total := dp[left][k-1] + coins + dp[k+1][right]
                
                if total > dp[left][right] {
                    dp[left][right] = total
                }
            }
        }
    }
    
    return dp[1][n]
}