func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 최대값, 최소값 및 결과 초기화
	maxProd, minProd, result := nums[0], nums[0], nums[0]

	// 배열의 두 번째 요소부터 순차적으로 탐색
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// 음수가 곱해지면 최대값과 최소값이 바뀔 수 있으므로 교환
		if num < 0 {
			maxProd, minProd = minProd, maxProd
		}

		// 현재 숫자와 곱한 값을 비교하여 최대값/최소값 갱신
		maxProd = max(num, maxProd*num)
		minProd = min(num, minProd*num)

		// 최대 곱을 결과에 갱신
		result = max(result, maxProd)
	}

	return result
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}