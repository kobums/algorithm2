func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }
    
    maxArea := 0
    rows := len(grid)
    cols := len(grid[0])
    
    // Helper function to perform DFS and calculate island area
    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        // 그리드의 범위를 벗어나거나 물(0)을 만나면 면적 0을 반환
        if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
            return 0
        }
        
        // 현재 위치를 방문 처리 (0으로 변경)
        grid[r][c] = 0
        
        // 상하좌우로 탐색하며 면적을 계산
        area := 1
        area += dfs(r+1, c) // 아래
        area += dfs(r-1, c) // 위
        area += dfs(r, c+1) // 오른쪽
        area += dfs(r, c-1) // 왼쪽
        
        return area
    }
    
    // 그리드를 순회하면서 섬의 최대 면적을 계산
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 1 {
                maxArea = max(maxArea, dfs(r, c))
            }
        }
    }
    
    return maxArea
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}