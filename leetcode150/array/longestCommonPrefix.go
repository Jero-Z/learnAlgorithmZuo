package array

/**
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。



示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/
func longestCommonPrefix(strs []string) string {
	var ans string
	for i := 0; i < len(strs[0]); i++ {
		point := true
		for j := 1; j < len(strs); j++ {

			if len(strs[j]) == 0 || len(strs[j]) < i || strs[j][i] != strs[0][i] {
				point = false
			}
		}
		if point {
			ans += string(strs[0][i])
		} else {
			break
		}

	}
	return ans
}
