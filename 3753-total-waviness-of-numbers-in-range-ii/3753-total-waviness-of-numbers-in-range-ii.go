// LeetCode 인터페이스: 매개변수와 반환 타입을 int64로 변경
func totalWaviness(num1 int64, num2 int64) int64 {
	return countUpTo(num2) - countUpTo(num1-1)
}

func countUpTo(num int64) int64 {
	if num < 100 {
		return 0
	}
	s := strconv.FormatInt(num, 10)
	n := len(s)

	// 다차원 슬라이스 캐싱 초기화 ([idx][prev1][prev2][isLess][isStart])
	// prev1, prev2는 0~9(미지정 시 10) 총 11가지 상태가 필요합니다.
	dp := make([][][][][]int64, n)
	for i := range dp {
		dp[i] = make([][][][]int64, 11)
		for j := range dp[i] {
			dp[i][j] = make([][][]int64, 11)
			for k := range dp[i][j] {
				dp[i][j][k] = make([][]int64, 2)
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = make([]int64, 2)
					for m := range dp[i][j][k][l] {
						dp[i][j][k][l][m] = -1
					}
				}
			}
		}
	}

	var dfs func(idx int, prev1, prev2 int, isLess, isStart bool) int64
	dfs = func(idx int, prev1, prev2 int, isLess, isStart bool) int64 {
		if idx == n {
			return 0
		}

		il := 0
		if isLess {
			il = 1
		}
		is := 0
		if isStart {
			is = 1
		}

		if dp[idx][prev1][prev2][il][is] != -1 {
			return dp[idx][prev1][prev2][il][is]
		}

		limit := 9
		if !isLess {
			limit = int(s[idx] - '0')
		}

		var res int64 = 0
		for i := 0; i <= limit; i++ {
			nextLess := isLess || (i < limit)
			nextStart := isStart && (i == 0)

			// 피크(Peak) 및 밸리(Valley) 체크
			var waviness int64 = 0
			if !isStart && prev1 != 10 && prev2 != 10 {
				if prev1 > prev2 && prev1 > i { // Peak
					waviness = 1
				} else if prev1 < prev2 && prev1 < i { // Valley
					waviness = 1
				}
			}

			var n1, n2 int
			if nextStart {
				n1, n2 = 10, 10
			} else {
				n1, n2 = i, prev1
			}

			// 현재 자릿수의 waviness 조합 수 계산 후 재귀 호출 누적
			combinations := countCombinations(n, idx+1, nextLess, s)
			res += waviness*combinations + dfs(idx+1, n1, n2, nextLess, nextStart)
		}

		dp[idx][prev1][prev2][il][is] = res
		return res
	}

	return dfs(0, 10, 10, false, true)
}

// 이후 자릿수에서 발생할 수 있는 유효 숫자 조합의 수를 반환하는 보조 함수 (int64 처리)
func countCombinations(length, idx int, isLess bool, s string) int64 {
	if idx == length {
		return 1
	}
	if isLess {
		var res int64 = 1
		for i := idx; i < length; i++ {
			res *= 10
		}
		return res
	}
	val, _ := strconv.ParseInt(s[idx:], 10, 64)
	return val + 1
}
