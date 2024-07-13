func groupAnagrams(strs []string) [][]string {
    tmp := make(map[int][]string, len(strs)/2)
    for _, v := range strs {
        c := encode(v)
        tmp[c] = append(tmp[c], v)
    }

    res := make([][]string, 0, len(tmp))
    for _, v := range tmp {
        res = append(res, v)
    }

    return res
}

var prime = []int{5, 71, 37, 29, 2, 53, 59, 19, 11, 83, 79, 31, 43, 13, 7, 67, 97, 23, 17, 3, 41, 73, 47, 89, 61, 101}


func encode(s string) int {
    res := 1
    for i := range s {
        res *= prime[s[i] - 'a']
    }

    return res
}