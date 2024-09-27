func findOrder(numCourses int, prerequisites [][]int) []int {
    // 그래프를 인접 리스트로 표현
    graph := make([][]int, numCourses)
    indegree := make([]int, numCourses)
    
    // 그래프 구성 및 indegree 계산
    for _, pre := range prerequisites {
        graph[pre[1]] = append(graph[pre[1]], pre[0])
        indegree[pre[0]]++
    }
    
    // Indegree가 0인 코스를 큐에 넣음
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 {
            queue = append(queue, i)
        }
    }
    
    // 위상 정렬 결과를 저장할 배열
    result := []int{}
    
    // 큐에서 코스를 하나씩 처리하며 Indegree를 업데이트
    for len(queue) > 0 {
        course := queue[0]
        queue = queue[1:]
        result = append(result, course)
        
        // 현재 코스를 선수 조건으로 참조하는 코스의 Indegree를 감소
        for _, next := range graph[course] {
            indegree[next]--
            if indegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }
    
    // 모든 코스를 처리했으면 결과를 반환, 그렇지 않으면 빈 배열 반환
    if len(result) == numCourses {
        return result
    }
    return []int{}
}