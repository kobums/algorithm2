func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	// 1. 땅 -> 물 순서로 탈 때의 최솟값 계산
	minLandEnd := math.MaxInt32
	for i := 0; i < len(landStartTime); i++ {
		end := landStartTime[i] + landDuration[i]
		if end < minLandEnd {
			minLandEnd = end
		}
	}

	ans1 := math.MaxInt32
	for j := 0; j < len(waterStartTime); j++ {
		// 땅이 끝난 시간과 물 놀이기구 시작 시간 중 더 늦은 시간에 시작
		start := minLandEnd
		if waterStartTime[j] > start {
			start = waterStartTime[j]
		}
		end := start + waterDuration[j]
		if end < ans1 {
			ans1 = end
		}
	}

	// 2. 물 -> 땅 순서로 탈 때의 최솟값 계산
	minWaterEnd := math.MaxInt32
	for j := 0; j < len(waterStartTime); j++ {
		end := waterStartTime[j] + waterDuration[j]
		if end < minWaterEnd {
			minWaterEnd = end
		}
	}

	ans2 := math.MaxInt32
	for i := 0; i < len(landStartTime); i++ {
		// 물이 끝난 시간과 땅 놀이기구 시작 시간 중 더 늦은 시간에 시작
		start := minWaterEnd
		if landStartTime[i] > start {
			start = landStartTime[i]
		}
		end := start + landDuration[i]
		if end < ans2 {
			ans2 = end
		}
	}

	// 두 경우 중 가장 이른 종료 시간 반환
	if ans1 < ans2 {
		return ans1
	}
	return ans2
}
