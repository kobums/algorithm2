func leftRightDifference(nums []int) []int {
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }

    ans := make([]int, len(nums))
    leftSum := 0
    
    for i, num := range nums {
        rightSum := totalSum - leftSum - num
        
        // 두 합의 절대값 계산
        diff := leftSum - rightSum
        if diff < 0 {
            diff = -diff
        }
        
        ans[i] = diff
        leftSum += num
    }
    
    return ans
}
