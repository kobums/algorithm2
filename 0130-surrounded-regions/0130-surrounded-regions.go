func solve(board [][]byte) {
    if len(board) == 0 || len(board[0]) == 0 {
        return
    }

    rows := len(board)
    cols := len(board[0])

    // 방향 벡터: 상, 하, 좌, 우
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    // DFS 함수
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // 보드를 벗어났거나 'O'가 아니면 종료
        if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != 'O' {
            return
        }

        // 현재 'O'를 임시로 'T'로 표시
        board[r][c] = 'T'

        // 상하좌우로 이동하여 연결된 'O'를 탐색
        for _, dir := range directions {
            dfs(r+dir[0], c+dir[1])
        }
    }

    // 1. 테두리에 있는 'O'들에 대해 DFS를 실행하여 'T'로 변환
    for r := 0; r < rows; r++ {
        if board[r][0] == 'O' {
            dfs(r, 0)
        }
        if board[r][cols-1] == 'O' {
            dfs(r, cols-1)
        }
    }
    for c := 0; c < cols; c++ {
        if board[0][c] == 'O' {
            dfs(0, c)
        }
        if board[rows-1][c] == 'O' {
            dfs(rows-1, c)
        }
    }

    // 2. 나머지 'O'는 'X'로 변환하고, 'T'는 다시 'O'로 변환
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if board[r][c] == 'O' {
                board[r][c] = 'X'
            } else if board[r][c] == 'T' {
                board[r][c] = 'O'
            }
        }
    }
}