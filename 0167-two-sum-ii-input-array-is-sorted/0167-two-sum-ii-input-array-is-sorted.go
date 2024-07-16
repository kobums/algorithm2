func twoSum(numbers []int, target int) []int {
    arr := make(map[int]int, len(numbers))

    for i, v := range numbers {
        if arr[target - v] != 0 {
            return []int{arr[target - v], i + 1}
        }
        arr[v] = i + 1
    }

    return nil
}