package array

/**
989. 数组形式的整数加法

整数的 数组形式  num 是按照从左到右的顺序表示其数字的数组。

例如，对于 num = 1321 ，数组形式是 [1,3,2,1] 。
给定 num ，整数的 数组形式 ，和整数 k ，返回 整数 num + k 的 数组形式 。



示例 1：

输入：num = [1,2,0,0], k = 34
输出：[1,2,3,4]
解释：1200 + 34 = 1234
示例 2：

输入：num = [2,7,4], k = 181
输出：[4,5,5]
解释：274 + 181 = 455
示例 3：

输入：num = [2,1,5], k = 806
输出：[1,0,2,1]
解释：215 + 806 = 1021
思路：

*/
func addToArrayForm(num []int, k int) []int {
	var ans []int
	var carry int
	i := len(num) - 1

	for i >= 0 || k != 0 {
		x, y := 0, 0
		if i >= 0 {
			x = num[i]
		}
		if k != 0 {
			y += k % 10
		}
		sum := x + y + carry
		ans = append(ans, sum%10)
		carry = sum / 10
		i--
		k /= 10
	}
	if carry != 0 {
		ans = append(ans, carry)
	}
	for l, r := 0, len(ans)-1; l < r; {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}
	return ans
}
