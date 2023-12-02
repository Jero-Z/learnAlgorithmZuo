package string

/**
541. 反转字符串 II
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。


示例 1：

输入：s = "abcdefg", k = 2
输出："bacdfeg"
示例 2：

输入：s = "abcd", k = 2
输出："bacd"

读题思路：
给定一个字符串和一个长度，反转2倍的长度区间的字符串
如果剩余字符串的长度小于k 反转全部；abcdef k=2 ； 2*2
如果剩余字符串的长度小于2*k 但是大于k就反转k个字符，其他的不变，否则将后面的
*/
func reverseStr(s string, k int) string {
	sliceStr := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if i+k <= length {
			// 反转前2k个字符串
			reverse(sliceStr[i : i+k])
		} else {
			//反转全部
			reverse(sliceStr[i:length])
		}
	}
	return string(sliceStr)
}
func reverse(b []byte) {
	for l, r := 0, len(b)-1; l < r; {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}
