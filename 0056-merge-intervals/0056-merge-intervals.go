func merge(intervals [][]int) [][]int {
    // 구간을 시작점을 기준으로 정렬합니다.
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    merged := [][]int{}
    
    for _, interval := range intervals {
        // merged가 비었거나, 현재 구간이 마지막 병합된 구간과 겹치지 않을 때
        if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
            merged = append(merged, interval)
        } else {
            // 겹치는 경우: 마지막 병합된 구간의 끝을 현재 구간의 끝과 비교하여 더 큰 값을 취함
            merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
        }
    }
    
    return merged
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}