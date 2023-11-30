package baseSort

/**
704. 二分查找

给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

*/
func search(nums []int, target int) int {
	if len(nums) == 0 || nums == nil {
		return -1
	}

	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//findLeft 寻找比给定目标大的数据,并返回其位置
func findLeft(nums []int, target int) int {
	if len(nums) == 0 || nums == nil {
		return -1
	}

	l, r, m := 0, len(nums)-1, 0
	ans := -1
	for l <= r {
		m = l + ((r - l) >> 1)
		if nums[m] >= target {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}

//findRight 寻找比给定目标小的最右位置的数据，并返回其位置
func findRight(nums []int, target int) int {
	if len(nums) == 0 || nums == nil {
		return -1
	}

	l, r, m := 0, len(nums)-1, 0
	ans := -1
	for l <= r {
		m = l + ((r - l) >> 1)
		if nums[m] <= target {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return ans
}

//findPeakElement 寻找峰值，并返回其所在位置
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[n-1] > nums[n-2] {
		return n - 1
	}

	l, r, m, ans := 1, n-2, 0, -1

	for l <= r {
		m = (l + r) / 2
		if nums[m-1] > nums[m] {
			r = m - 1
		} else if nums[m] < nums[m+1] {
			l = m + 1
		} else {
			ans = m
			break
		}
	}
	return ans
}
