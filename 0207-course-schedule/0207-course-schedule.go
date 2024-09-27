func canFinish(numCourses int, prerequisites [][]int) bool {
    // 그래프를 인접 리스트로 표현
    graph := make([][]int, numCourses)
    for _, pre := range prerequisites {
        graph[pre[1]] = append(graph[pre[1]], pre[0])
    }

    // 방문 상태 배열: 0은 미방문, 1은 방문 중, 2는 방문 완료
    visited := make([]int, numCourses)

    // DFS로 사이클 감지
    var dfs func(node int) bool
    dfs = func(node int) bool {
        if visited[node] == 1 { // 방문 중인 노드를 다시 방문하면 사이클 발생
            return false
        }
        if visited[node] == 2 { // 이미 방문 완료한 노드는 다시 탐색할 필요 없음
            return true
        }

        // 방문 중으로 표시
        visited[node] = 1

        // 인접 노드들 탐색
        for _, neighbor := range graph[node] {
            if !dfs(neighbor) {
                return false
            }
        }

        // 방문 완료로 표시
        visited[node] = 2
        return true
    }

    // 모든 노드에 대해 DFS 수행
    for i := 0; i < numCourses; i++ {
        if !dfs(i) {
            return false
        }
    }

    return true
}