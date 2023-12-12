package baseSort

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	m := len(nums) / 2
	left := mergeSort(nums[:m])
	right := mergeSort(nums[m:])

	return merge(left, right)

}

func merge(left, right []int) []int {
	// 创建一个合适的容量
	res := make([]int, len(left)+len(right))
	// i = left index
	// j = right index
	// k = res index
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			// 左边段小
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}
	// 剩余元素中左边和右边最少&最多有一个越界的情况存在
	//将剩余的数据再进行写入
	for i < len(left) {
		res[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		res[k] = right[j]
		j++
		k++
	}
	return res
}
