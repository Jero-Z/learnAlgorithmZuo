package string

/**
28. 找出字符串中第一个匹配项的下标

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
如果 needle 不是 haystack 的一部分，则返回  -1 。

示例 1：

输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
示例 2：

输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1
思路：
循环范围haystack- len(needle)
思路：
先判定查询的字符串长度是否小于给定字符串长度
进阶：
KPM
*/
func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(haystack) < len(needle) || len(needle) == 0 {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		op := ""
		if haystack[i] != needle[0] {
			continue // 若不等于直接跳出首字母的位置|查找需要查找的首字母的起点
		}
		if i+len(needle) > len(haystack) { // 防止越界
			return -1
		}

		for j := 0; j < len(needle); j++ {

			op += string(haystack[i+j])
		}
		if needle == op {
			return i
		}
	}
	return -1
}
