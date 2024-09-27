var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func pacificAtlantic(heights [][]int) [][]int {
    if len(heights) == 0 || len(heights[0]) == 0 {
        return [][]int{}
    }
    
    rows := len(heights)
    cols := len(heights[0])
    
    pacific := make([][]bool, rows)
    atlantic := make([][]bool, rows)
    for i := range heights {
        pacific[i] = make([]bool, cols)
        atlantic[i] = make([]bool, cols)
    }
    
    var dfs func(r, c int, ocean [][]bool)
    dfs = func(r, c int, ocean [][]bool) {
        ocean[r][c] = true
        
        for _, d := range directions {
            nr, nc := r+d[0], c+d[1]
            if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !ocean[nr][nc] && heights[nr][nc] >= heights[r][c] {
                dfs(nr, nc, ocean)
            }
        }
    }
    
    // 태평양: 왼쪽 및 위쪽 가장자리에서 시작
    for r := 0; r < rows; r++ {
        dfs(r, 0, pacific)
    }
    for c := 0; c < cols; c++ {
        dfs(0, c, pacific)
    }
    
    // 대서양: 오른쪽 및 아래쪽 가장자리에서 시작
    for r := 0; r < rows; r++ {
        dfs(r, cols-1, atlantic)
    }
    for c := 0; c < cols; c++ {
        dfs(rows-1, c, atlantic)
    }
    
    result := [][]int{}
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if pacific[r][c] && atlantic[r][c] {
                result = append(result, []int{r, c})
            }
        }
    }
    
    return result
}