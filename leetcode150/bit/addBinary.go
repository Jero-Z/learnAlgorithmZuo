package bit

/**
67. 二进制求和

给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。

输入:a = "11", b = "1"
输出："100"
示例 2：

输入：a = "1010", b = "1011"
输出："10101"

思路：
双指针 + 进位标识

*/
func addBinary(a string, b string) string {
	if a == "0" {
		return b
	}
	if b == "0" {
		return a
	}



	ai, bi := len(a), len(b)
	var temp uint8
	for ai > 0 || bi > 0 || temp != 0 {
		if ai < len(a) && bi < len(b) {

		}
	}
	return ""
}
