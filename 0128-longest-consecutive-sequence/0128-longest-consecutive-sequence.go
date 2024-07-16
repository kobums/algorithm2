func longestConsecutive(nums []int) int {
    arr := make(map[int]int)
    for _, v := range nums {
        arr[v] = 1
    }

    keys := make([]int, 0, len(arr))
    for i, k := range arr {
		fmt.Println(k)
        keys = append(keys, i)
    }

    sort.Ints(keys)
    fmt.Println(keys)

    l := 0
    max := 0
    for i, k := range keys {
        if i == 0 {
            l = 1
            max = 1
            continue
        }

        if k-1 == keys[i-1] {
            l += 1
        } else {
			l = 1
		}

        if l > max {
            max = l
        }
    }
    return max
}