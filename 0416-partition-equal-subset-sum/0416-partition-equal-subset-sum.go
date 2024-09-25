func canPartition(nums []int) bool {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    
    // 합이 홀수면 두 부분으로 나눌 수 없으므로 false
    if sum%2 != 0 {
        return false
    }
    
    target := sum / 2
    dp := make([]bool, target+1)
    dp[0] = true // 0을 만드는 부분 집합은 항상 존재
    
    // 각 숫자에 대해 dp 테이블 갱신
    for _, num := range nums {
        // 뒤에서부터 순회하며 dp 갱신 (중복 계산 방지)
        for j := target; j >= num; j-- {
            dp[j] = dp[j] || dp[j-num]
        }
    }
    
    return dp[target]
}
