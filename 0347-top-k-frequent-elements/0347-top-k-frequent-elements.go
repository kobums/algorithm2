func topKFrequent(nums []int, k int) []int {
    res := make([]int, 0, k)
    arr := make(map[int]int, len(nums))

    for _, v := range nums {
        arr[v] += 1
    }

    count := make([]int, 0, len(arr))
    for _, v := range arr {
        count = append(count, v)
    }

    sort.Ints(count)

    min := count[len(count)-k]

    for i, v := range arr {
        if v >= min {
            res = append(res, i)
        }
    }

    return res
}