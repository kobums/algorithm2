func jump(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    
    jumps, curEnd, farthest := 0, 0, 0
    
    for i := 0; i < len(nums)-1; i++ {
        farthest = max(farthest, i+nums[i]) // 현재 위치에서 도달할 수 있는 가장 먼 곳
        
        if i == curEnd { // 현재 범위의 끝에 도달했을 때
            jumps++         // 점프 증가
            curEnd = farthest // 다음 범위 설정
        }
    }
    
    return jumps
}

// 두 정수 중 큰 값을 반환하는 유틸리티 함수
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}