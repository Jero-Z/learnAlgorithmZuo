package dp

/**
42. 接雨水

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9
*/

func trap(height []int) int {
	leftIndex, rightIndex := 0, len(height)-1
	var leftMax, rightMax, ans int

	for leftIndex < rightIndex { // 开始循环
		if height[leftIndex] < height[rightIndex] {
			if height[leftIndex] > leftMax {
				leftMax = height[leftIndex]
			} else {
				ans += leftMax - height[leftIndex]
			}
			leftIndex++
		} else {
			if height[rightIndex] > rightMax {
				rightMax = height[rightIndex]
			} else {
				ans += rightMax - height[rightIndex]
			}
			rightIndex--
		}
	}
	return ans
}
