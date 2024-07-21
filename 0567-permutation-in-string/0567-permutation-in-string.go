func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    var s1count, s2count [26]int

    for i, _ := range s1 {
        s1count[s1[i]-'a']++
        s2count[s2[i]-'a']++
        fmt.Println(s1count)
        fmt.Println(s2count)
    }

    for i := len(s1); i < len(s2); i++ {
        if s1count == s2count {
            return true
        }
        s2count[s2[i]-'a']++
        s2count[s2[i-len(s1)]-'a']--
    }

    return s1count == s2count
}