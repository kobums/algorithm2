func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{}
    combination := []int{}
    backtracking(candidates, target, combination, 0, &result)
    return result
}

func backtracking(candidates []int, target int, combination []int, start int, result *[][]int) {
    if target == 0 {
        temp := make([]int, len(combination))
        copy(temp, combination)
        *result = append(*result, temp)
        return
    }
    if target < 0 {
        return
    }

    for i := start; i < len(candidates); i++ {
        combination = append(combination, candidates[i])
        backtracking(candidates, target-candidates[i], combination, i, result)
        combination = combination[:len(combination)-1]
    }
}