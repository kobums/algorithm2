func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    
    numIslands := 0
    rows := len(grid)
    cols := len(grid[0])
    
    // Helper function to perform DFS
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // 그리드의 범위를 벗어나거나 물(0)을 만나면 종료
        if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
            return
        }
        
        // 현재 위치를 방문 처리 (0으로 변경)
        grid[r][c] = '0'
        
        // 상하좌우로 계속 탐색
        dfs(r+1, c) // 아래쪽
        dfs(r-1, c) // 위쪽
        dfs(r, c+1) // 오른쪽
        dfs(r, c-1) // 왼쪽
    }
    
    // 그리드를 순회하면서 섬을 찾음
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == '1' { // 새로운 섬을 찾으면
                numIslands++
                dfs(r, c) // 그 섬을 DFS로 모두 탐색하여 방문 처리
            }
        }
    }
    
    return numIslands
}