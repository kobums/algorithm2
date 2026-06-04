func totalWaviness(num1 int, num2 int) int {
    // Range I은 범위가 작으므로 굳이 전역 Prefix Sum을 쓰지 않고 
    // 반복문으로 직관적으로 합산해도 시간 초과(TLE) 없이 통과합니다.
    total := 0
    for i := num1; i <= num2; i++ {
        total += countWaviness(i)
    }
    return total
}

func countWaviness(n int) int {
    if n < 100 {
        return 0
    }
    
    // 자릿수 추출 (1의 자리부터 역순으로 배열에 저장됨)
    var digits []int
    temp := n
    for temp > 0 {
        digits = append(digits, temp%10)
        temp /= 10
    }

    c := 0
    // 첫 자릿수와 마지막 자릿수를 제외하고 내부 피크/밸리 검사
    for i := 1; i < len(digits)-1; i++ {
        prev, curr, next := digits[i+1], digits[i], digits[i-1]
        
        // Peak 조건
        if curr > prev && curr > next {
            c++
        }
        // Valley 조건
        if curr < prev && curr < next {
            c++
        }
    }
    return c
}
