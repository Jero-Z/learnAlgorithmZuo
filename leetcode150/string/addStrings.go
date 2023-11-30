package string

import "strconv"

/**
415. 字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。

你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
示例 1：

输入：num1 = "11", num2 = "123"
输出："134"
示例 2：

输入：num1 = "456", num2 = "77"
输出："533"
示例 3：

输入：num1 = "0", num2 = "0"
输出："0"
双指针同时对两个字符串进行处理
先获取两个字符串的长度
进行循环处理相加，记得处理进位

*/
func addStrings(num1 string, num2 string) string {

	carry := 0
	ans := ""

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		i--
		j--
		res := x + y + carry
		ans = strconv.Itoa(res%10) + ans
		carry = res / 10
	}

	return ans
}
