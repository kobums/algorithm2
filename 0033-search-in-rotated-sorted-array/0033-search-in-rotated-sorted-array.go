func search(nums []int, target int) int {
    rotated := indexOfMin(nums)
    size := len(nums)
    left, right := 0, size-1

    for left <= right {
        mid := (left + right) / 2
        rotatedMid := (rotated + mid) % size
        switch {
            case nums[rotatedMid] < target:
                left = mid + 1
            case nums[rotatedMid] > target:
                right = mid - 1
            default:
                return rotatedMid
        }
    }
    return -1
}

func indexOfMin(nums []int) int {
    size := len(nums)
    left, right := 0, size-1

    for left < right {
        mid := (left + right) / 2
        if nums[right] < nums[mid] {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return left
}