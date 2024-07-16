func productExceptSelf(nums []int) []int {
    size := len(nums)

    res := make([]int, size)

    left := 1
    for i, v := range nums {
        res[i] = left
        left *= v
    }

    right := 1
    for i := size - 1; i >= 0; i-- {
        res[i] *= right
        right *= nums[i]
    }

    return res
}