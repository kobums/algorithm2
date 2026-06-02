// int 타입의 최솟값을 구하는 헬퍼 함수
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// int 타입의 최댓값을 구하는 헬퍼 함수
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	ans := math.MaxInt32

	// 플랜 1: 육상(Land) 놀이기구를 먼저 타고, 수상(Water) 놀이기구를 타는 경우
	for i := 0; i < len(landStartTime); i++ {
		landFinish := landStartTime[i] + landDuration[i]
		for j := 0; j < len(waterStartTime); j++ {
			// 수상 놀이기구는 자신의 시작 시간과 육상 놀이기구가 끝난 시간 중 더 늦은 시간에 시작할 수 있습니다.
			waterStart := max(waterStartTime[j], landFinish)
			finish := waterStart + waterDuration[j]
			ans = min(ans, finish)
		}
	}

	// 플랜 2: 수상(Water) 놀이기구를 먼저 타고, 육상(Land) 놀이기구를 타는 경우
	for i := 0; i < len(waterStartTime); i++ {
		waterFinish := waterStartTime[i] + waterDuration[i]
		for j := 0; j < len(landStartTime); j++ {
			// 육상 놀이기구는 자신의 시작 시간과 수상 놀이기구가 끝난 시간 중 더 늦은 시간에 시작할 수 있습니다.
			landStart := max(landStartTime[j], waterFinish)
			finish := landStart + landDuration[j]
			ans = min(ans, finish)
		}
	}

	return ans
}
