func longestPalindrome(s string) string {
    n := len(s)
    if n < 2 {
        return s
    }

    // DP 테이블 생성
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }

    start, maxLen := 0, 1

    // 길이 1 초기화
    for i := 0; i < n; i++ {
        dp[i][i] = true
    }

    // 길이(len)를 기준으로 반복
    for length := 2; length <= n; length++ {
        for i := 0; i <= n-length; i++ {
            j := i + length - 1
            
            if s[i] == s[j] {
                // 길이가 2이거나 안쪽이 팰린드롬이면
                if length == 2 || dp[i+1][j-1] {
                    dp[i][j] = true
                    if length > maxLen {
                        start = i
                        maxLen = length
                    }
                }
            }
        }
    }

    return s[start : start+maxLen]
}