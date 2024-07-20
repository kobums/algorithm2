func lengthOfLongestSubstring(s string) int {
    location := make([]int, 256)
    for i := range location {
		location[i] = -1
	}
    max, left := 0, 0
    

    for i, _ := range s {
        if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > max {
			max = i + 1 - left
		}
		location[s[i]] = i
    }

    return max
}