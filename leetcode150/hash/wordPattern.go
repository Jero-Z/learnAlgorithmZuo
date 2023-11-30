package hash

import "strings"

/**
290. 单词规律

给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。

思路：
分割字符串成为slice

建立pattern 和每个单词的映射表
建立每个单词 和每个byte的映射表

如果映射表中不存在则写入
如果映射表中的存在则进行判断是否和映射表中的一致
dog=>a
cat=>b

*/
func wordPattern(pattern string, s string) bool {
	strNums := strings.Split(s, " ")
	if len(pattern) != len(strNums) {
		return false
	}
	wordMap := map[string]byte{}
	patternMap := map[byte]string{}

	for i, num := range strNums {
		ch := pattern[i] // 当前位置对应的byte

		if wordMap[num] > 0 && wordMap[num] != ch ||
			patternMap[ch] != "" && patternMap[ch] != num {
			return false
		}
		wordMap[num] = ch
		patternMap[ch] = num

	}

	return true
}
