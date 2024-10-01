func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    // 첫 번째 행과 첫 번째 열은 경로가 1개씩 있음
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }

    // 나머지 dp 테이블 채우기
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }

    // 오른쪽 아래 코너에 도달할 수 있는 경로의 수 반환
    return dp[m-1][n-1]
}