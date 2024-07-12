func twoSum(nums []int, target int) []int {
    index := make(map[int]int, len(nums))
    for i, v := range nums {
        if j, ok := index[target-v]; ok {
            return []int{i, j}
        }
        index[v] = i
    }
    return nil
}