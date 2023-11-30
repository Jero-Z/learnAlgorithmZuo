package array

/**
169. 多数元素

给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
示例 1：

输入：nums = [3,2,3]
输出：3
示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2

思路：
假设有一个城堡被占领了，占领的士兵数量为n
跟我方相同的则增加一个士兵
遇见一个和我方不同的则减一个士兵
若士兵被减少为0的时候城堡易主
因为多数元素是指在数组中出现的次数大于[n/2] ，所以必然有一个胜者

*/
func majorityElement(nums []int) int {
	if len(nums) <= 2 { // 数组长度小于等于2 就可以返回第一个元素
		return nums[0]
	}
	winner := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == winner {
			count++
		} else {

			count--
			if count == 0 {
				winner = nums[i]
				count++
			}
		}
	}
	return winner
}
func majorityElement2(nums []int) int {
	if len(nums) <= 2 { // 数组长度小于等于2 就可以返回第一个元素
		return nums[0]
	}
	winner := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == winner {
			count++
		} else if count == 0 {
			winner = nums[i]
			count++
		} else {
			count--
		}
	}
	return winner
}
func majorityElement1(nums []int) int {
	if len(nums) <= 2 { // 数组长度小于等于2 就可以返回第一个元素
		return nums[0]
	}
	counter := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		counter[nums[i]] += 1
	}
	var max, r int

	for c, i := range counter {

		if i > max {
			r = c
			max = i
		}
	}
	return r
}
