func combinationSum2(candidates []int, target int) [][]int {
    result := [][]int{}
    connect := []int{}
    sort.Ints(candidates)
    backtrack(candidates, 0, target, connect, &result)
    return result
}

func backtrack(candidates []int, order, target int, connect []int, result *[][]int) {
    if target == 0 {
        temp := make([]int, len(connect))
        copy(temp, connect)
        *result = append(*result, temp)
        return
    }

    for i := order; i < len(candidates); i++ {
        if candidates[i] > target {
            continue
        }
        if i > order && candidates[i] == candidates[i-1] {
			continue
		}
        connect = append(connect, candidates[i])
        backtrack(candidates, i+1, target - candidates[i], connect, result)
        connect = connect[:len(connect) -1]
    }
}