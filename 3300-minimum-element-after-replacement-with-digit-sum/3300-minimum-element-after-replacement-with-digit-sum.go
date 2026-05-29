func minElement(nums []int) int {
	minVal := 1000 // 초기값 설정 (문제 제약사항보다 큰 수)

	for _, num := range nums {
		// 각 숫자의 자릿수 합 계산
		sum := 0
		temp := num
		for temp > 0 {
			sum += temp % 10
			temp /= 10
		}
		
		// 최솟값 업데이트
		if sum < minVal {
			minVal = sum
		}
	}

	return minVal
}
