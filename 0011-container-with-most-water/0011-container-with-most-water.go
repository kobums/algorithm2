func maxArea(height []int) int {
    i, j := 0, len(height) -1
    res := 0

    for i < j {
        a, b := height[i], height[j]

        h := min(a, b)

         
        area := h * (j-i) 
        if res < area {
            res = area
        }

        if a < b {
            i++
        } else {
            j--
        }
    }

    return res
}


func min (i, j int) int {
    if i < j {
        return i
    }
    return j
}