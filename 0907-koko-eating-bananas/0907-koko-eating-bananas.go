func minEatingSpeed(piles []int, h int) int {
    mx, sum := 0, 0

    for _, v := range piles {
        sum += v
        mx = max(mx, v)
    }

    aver := (sum + h - 1) / h

    for aver < mx {
        m := (aver + mx) / 2
        if howManyEat(m, h, piles) {
            mx = m
        } else {
            aver = m + 1
        }
    }
    
    return aver
}

func howManyEat(count int, h int, piles []int) bool {
    r := 1 / float64(count)
    for _, v := range piles {
        fv := float64(v)
        h -= int(math.Ceil(fv * r))
        if h < 0 {
            return false
        }
    }
    return true
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}