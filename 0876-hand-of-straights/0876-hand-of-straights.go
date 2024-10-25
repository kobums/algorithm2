func isNStraightHand(hand []int, groupSize int) bool {
    // hand 길이가 groupSize로 나누어 떨어지지 않으면 바로 false 반환
    if len(hand) % groupSize != 0 {
        return false
    }

    // 카드를 정렬하여 작은 숫자부터 처리
    sort.Ints(hand)
    
    // 카드의 빈도를 저장할 map
    freq := make(map[int]int)
    
    // 각 카드의 빈도를 기록
    for _, card := range hand {
        freq[card]++
    }

    // 작은 숫자부터 처리
    for _, card := range hand {
        // 이미 처리된 카드라면 무시
        if freq[card] == 0 {
            continue
        }

        // 현재 카드부터 groupSize 만큼의 연속된 숫자가 있는지 확인
        for i := 0; i < groupSize; i++ {
            if freq[card + i] <= 0 {
                return false
            }
            // 해당 카드를 사용했으므로 빈도를 감소
            freq[card + i]--
        }
    }
    
    return true
}