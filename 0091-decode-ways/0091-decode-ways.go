func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	// dp[i]는 s[0:i]까지 해독할 수 있는 방법의 수
	dp := make([]int, n+1)
	dp[0] = 1 // 빈 문자열은 1가지 방법으로 해독될 수 있음
	dp[1] = 1 // 첫 번째 문자가 '0'이 아닌 경우에만 1가지 방법

	for i := 2; i <= n; i++ {
		// 한 자리 숫자 확인
		oneDigit, _ := strconv.Atoi(s[i-1 : i])
		if oneDigit >= 1 && oneDigit <= 9 {
			dp[i] += dp[i-1]
		}

		// 두 자리 숫자 확인
		twoDigit, _ := strconv.Atoi(s[i-2 : i])
		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}