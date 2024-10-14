func maxSubArray(nums []int) int {
    // 배열의 첫 번째 원소로 최대 합과 현재 부분 배열의 합을 초기화
    maxSum := nums[0]
    currentSum := nums[0]
    
    // 두 번째 원소부터 순차적으로 탐색
    for i := 1; i < len(nums); i++ {
        // 현재 원소를 더하는 것이 유리한지, 현재 원소부터 새로 시작하는 것이 유리한지 선택
        currentSum = max(nums[i], currentSum + nums[i])
        // 최대 합을 갱신
        maxSum = max(maxSum, currentSum)
    }
    
    return maxSum
}

// 두 정수 중 더 큰 값을 반환하는 유틸리티 함수
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}