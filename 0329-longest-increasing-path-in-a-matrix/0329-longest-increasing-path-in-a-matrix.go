func longestIncreasingPath(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
    }
    
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    
    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        if memo[r][c] != 0 {
            return memo[r][c]
        }
        
        maxLen := 1
        for _, d := range dirs {
            nr, nc := r+d[0], c+d[1]
            
            if nr >= 0 && nr < m && nc >= 0 && nc < n &&
                matrix[nr][nc] > matrix[r][c] {
                if length := 1 + dfs(nr, nc); length > maxLen {
                    maxLen = length
                }
            }
        }
        
        memo[r][c] = maxLen
        return maxLen
    }
    
    result := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if length := dfs(i, j); length > result {
                result = length
            }
        }
    }
    
    return result
}