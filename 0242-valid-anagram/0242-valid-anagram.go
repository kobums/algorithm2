func isAnagram(s string, t string) bool {
    s1 := strings.Split(s, "")
    t1 := strings.Split(t, "")

    if len(s1) != len(t1) {
        return false
    }

    sort.Slice(s1, func(i, j int) bool {
        return s1[i] < s1[j]
    })

    sort.Slice(t1, func(i, j int) bool {
        return t1[i] < t1[j]
    })

    for i, _ := range s1 {
        if s1[i] != t1[i] {
            return false
        }
    }

    return true
}