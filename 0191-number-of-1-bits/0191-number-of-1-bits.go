func hammingWeight(n int) int {
    count := 0
    for n != 0 {
        n &= (n - 1) // 가장 오른쪽 1을 제거
        count++
    }
    return count
}