func permute(nums []int) [][]int {
     result := [][]int{}
    backtrack(nums, nil, &result)
    return result
}

// backtrack function generates permutations using backtracking.
func backtrack(nums []int, current []int, result *[][]int) {
    if len(nums) == 0 {
        temp := make([]int, len(current))
        copy(temp, current)
        *result = append(*result, temp)
        return
    }

    for i := 0; i < len(nums); i++ {
        newNums := append([]int{}, nums[:i]...)
        newNums = append(newNums, nums[i+1:]...)
        current = append(current, nums[i])
        backtrack(newNums, current, result)
        current = current[:len(current)-1]
    }
}