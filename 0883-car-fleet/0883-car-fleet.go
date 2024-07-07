func carFleet(target int, position []int, speed []int) int {
    size := len(position)
    rs := make([]recode, size)
    for i := range rs {
        p, s := position[i], speed[i]
        rs[i] = recode{
            initpos: p,
            arrivalTime: float64(target - p) / float64(s),
        }
    }

    sort.Slice(rs, func(i int, j int) bool{
        return rs[i].initpos > rs[j].initpos
    })

    feetTime := 0.
    res := 0

    for _, r := range rs {
        at := r.arrivalTime
        if feetTime < at {
            feetTime = at
            res++
        }
    }

    return res
}

type recode struct {
    initpos int
    arrivalTime float64
}