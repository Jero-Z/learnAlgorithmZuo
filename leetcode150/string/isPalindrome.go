package string

import "strings"

/*
125. 验证回文串

如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

字母和数字都属于字母数字字符。

给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
思路：
*/
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToUpper(s)
	head, tail := 0, len(s)-1
	for head < tail {
		if !isValidByte(s[head]) {
			head++
			continue
		}
		if !isValidByte(s[tail]) {
			tail--
			continue
		}
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}
func isValidByte(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
