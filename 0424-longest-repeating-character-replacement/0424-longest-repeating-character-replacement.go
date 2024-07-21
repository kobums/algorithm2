func characterReplacement(s string, k int) int {
    max, left := 0, 0

    count := make([]int, 256)

    for i, v := range s {
        count[v]++
        max = Max(max, count[v])

        if i+1 - left - max > k {
            count[s[left]]--
            left++
        }
    }

    return len(s) - left
}

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}