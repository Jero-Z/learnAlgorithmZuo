package array

/**
189. 轮转数组

给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
思路： k=3
先反转数组
[7,6,5,4,3,2,1]
再反转k位置的前面的元素
[5，6，7，4，3，2，1]
再反转k位置后面的元素
[5,6,7,1,2,3,4]
*/
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func reverse(a []int) {
	// 使用双指针，循环n/2次，进行两两调换
	for i, totalCount := 0, len(a); i < totalCount/2; i++ {
		a[i], a[totalCount-1-i] = a[totalCount-1-i], a[i]
	}
}
