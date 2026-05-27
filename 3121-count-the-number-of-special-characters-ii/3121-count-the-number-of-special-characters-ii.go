func numberOfSpecialChars(word string) int {
	// 알파벳 26개에 대해 첫 번째 대문자와 마지막 소문자의 인덱스를 저장할 배열
	// Go에서는 배열 선언 시 자동으로 0으로 초기화되므로, 0번 인덱스와의 혼동을 피하기 위해 -1로 초기화합니다.
	firstUpper := make([]int, 26)
	lastLower := make([]int, 26)
	for i := 0; i < 26; i++ {
		firstUpper[i] = -1
		lastLower[i] = -1
	}

	// 1. 문자열을 순회하며 각 문자의 인덱스를 기록
	for i := 0; i < len(word); i++ {
		char := word[i]
		if char >= 'a' && char <= 'z' {
			idx := char - 'a'
			lastLower[idx] = i // 마지막 소문자 위치 갱신
		} else if char >= 'A' && char <= 'Z' {
			idx := char - 'A'
			if firstUpper[idx] == -1 {
				firstUpper[idx] = i // 첫 번째 대문자 위치만 기록
			}
		}
	}

	specialCount := 0

	// 2. 두 배열을 비교하여 조건(모든 소문자가 첫 대문자보다 앞에 위치)을 만족하는지 확인
	for i := 0; i < 26; i++ {
		if lastLower[i] != -1 && firstUpper[i] != -1 {
			if lastLower[i] < firstUpper[i] {
				specialCount++
			}
		}
	}

	return specialCount
}
