func canJump(nums []int) bool {
    maxReach := 0 // 현재까지 도달할 수 있는 가장 먼 인덱스

    for i := 0; i < len(nums); i++ {
        if i > maxReach {
            // 현재 인덱스에 도달할 수 없다면, false 반환
            return false
        }
        // 현재 인덱스에서 도달할 수 있는 최대 거리를 계산
        maxReach = max(maxReach, i + nums[i])
    }
    
    return true // 마지막 인덱스까지 도달할 수 있다면 true 반환
}

// 두 정수 중 큰 값을 반환하는 유틸리티 함수
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}