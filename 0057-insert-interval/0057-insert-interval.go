func insert(intervals [][]int, newInterval []int) [][]int {
    result := [][]int{}
    i := 0
    n := len(intervals)

    // Step 1: Add all intervals that come before the newInterval
    for i < n && intervals[i][1] < newInterval[0] {
        result = append(result, intervals[i])
        i++
    }

    // Step 2: Merge overlapping intervals
    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])  // 기존의 끝과 병합한 끝을 비교하여 큰 값을 저장
        i++
    }
    result = append(result, newInterval)

    // Step 3: Add remaining intervals
    for i < n {
        result = append(result, intervals[i])
        i++
    }

    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}