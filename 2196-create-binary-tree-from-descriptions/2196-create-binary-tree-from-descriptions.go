func createBinaryTree(descriptions [][]int) *TreeNode {
    // 노드 값과 TreeNode 포인터를 매핑할 맵
    nodes := make(map[int]*TreeNode)
    // 자식 노드들의 값을 저장할 셋 (루트 노드 판별용)
    hasParent := make(map[int]bool)
    
    for _, desc := range descriptions {
        parentVal, childVal, isLeft := desc[0], desc[1], desc[2]
        
        // 1. 부모 노드 처리
        if _, exists := nodes[parentVal]; !exists {
            nodes[parentVal] = &TreeNode{Val: parentVal}
        }
        parent := nodes[parentVal]
        
        // 2. 자식 노드 처리
        if _, exists := nodes[childVal]; !exists {
            nodes[childVal] = &TreeNode{Val: childVal}
        }
        child := nodes[childVal]
        
        // 3. 자식을 부모에 연결
        if isLeft == 1 {
            parent.Left = child
        } else {
            parent.Right = child
        }
        
        // 4. 자식은 부모가 있음을 기록
        hasParent[childVal] = true
    }
    
    // 5. 부모가 없는 노드를 찾아 루트로 반환
    for _, desc := range descriptions {
        parentVal := desc[0]
        if !hasParent[parentVal] {
            return nodes[parentVal]
        }
    }
    
    return nil
}
