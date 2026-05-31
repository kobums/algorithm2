func asteroidsDestroyed(mass int, asteroids []int) bool {
	// 소행성 질량을 오름차순으로 정렬
	sort.Ints(asteroids)
	
	currentMass := int64(mass)
	
	for _, asteroid := range asteroids {
		// 현재 행성의 질량이 소행성 질량보다 작으면 파괴 불가
		if currentMass < int64(asteroid) {
			return false
		}
		// 소행성을 파괴하고 질량을 흡수
		currentMass += int64(asteroid)
	}
	
	return true
}
