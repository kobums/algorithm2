// 방향 벡터 (상, 하, 좌, 우)
var directions = [][]int{
    {1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func orangesRotting(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])
    freshOranges := 0
    queue := [][]int{}

    // 처음 상태에서 썩은 오렌지를 큐에 넣고, 신선한 오렌지 개수를 셉니다.
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 2 {
                queue = append(queue, []int{r, c})
            } else if grid[r][c] == 1 {
                freshOranges++
            }
        }
    }

    // 신선한 오렌지가 없으면 이미 모두 썩은 상태
    if freshOranges == 0 {
        return 0
    }

    minutes := 0

    // BFS 시작
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            current := queue[0]
            queue = queue[1:]

            r, c := current[0], current[1]
            // 상하좌우 탐색
            for _, direction := range directions {
                nr, nc := r+direction[0], c+direction[1]
                // 유효한 위치이면서 신선한 오렌지일 경우
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
                    // 신선한 오렌지를 썩게 만들고 큐에 추가
                    grid[nr][nc] = 2
                    queue = append(queue, []int{nr, nc})
                    freshOranges--
                }
            }
        }
        // 한 번의 탐색이 끝나면 분 증가
        minutes++
    }

    // BFS가 끝났을 때 신선한 오렌지가 남아 있으면 -1 반환
    if freshOranges > 0 {
        return -1
    }
    
    return minutes - 1
}