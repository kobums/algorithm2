func findMin(nums []int) int {
    l, r := 1, len(nums)

    for l < r && nums[l-1] < nums[l] {
        l++
    }

    return nums[l%r]
}