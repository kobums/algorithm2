func eraseOverlapIntervals(intervals [][]int) int {
    // 구간을 끝값을 기준으로 정렬
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    
    // 겹치는 구간을 카운트하기 위한 변수
    count := 0
    prevEnd := intervals[0][1]
    
    // 1번 구간부터 탐색 시작
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] < prevEnd {
            // 현재 구간이 이전 구간과 겹치면 제거해야 함
            count++
        } else {
            // 겹치지 않으면 이전 구간의 끝값을 업데이트
            prevEnd = intervals[i][1]
        }
    }
    
    return count
}