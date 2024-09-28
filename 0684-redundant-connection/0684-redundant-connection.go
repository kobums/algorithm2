func findRedundantConnection(edges [][]int) []int {
    parent := make([]int, len(edges)+1)
    rank := make([]int, len(edges)+1)

    // 초기화: 각 노드의 부모를 자신으로 설정
    for i := 1; i <= len(edges); i++ {
        parent[i] = i
    }

    // Find 함수: 경로 압축 기법 사용
    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }

    // Union 함수: union by rank 사용
    union := func(x, y int) bool {
        rootX := find(x)
        rootY := find(y)
        if rootX == rootY {
            return false
        }
        if rank[rootX] > rank[rootY] {
            parent[rootY] = rootX
        } else if rank[rootX] < rank[rootY] {
            parent[rootX] = rootY
        } else {
            parent[rootY] = rootX
            rank[rootX]++
        }
        return true
    }

    // 간선들을 처리하며, 사이클을 형성하는 간선을 찾음
    for _, edge := range edges {
        if !union(edge[0], edge[1]) {
            return edge
        }
    }
    return nil
}