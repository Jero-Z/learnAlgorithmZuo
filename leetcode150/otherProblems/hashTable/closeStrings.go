package hashTable

import (
	"reflect"
	"sort"
)

/**
1657. 确定两个字符串是否接近
如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：

操作 1：交换任意两个 现有 字符。
例如，abcde -> aecdb
操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。
例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）
你可以根据需要对任意一个字符串多次使用这两种操作。

给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。



示例 1：

输入：word1 = "abc", word2 = "bca"
输出：true
解释：2 次操作从 word1 获得 word2 。
执行操作 1："abc" -> "acb"
执行操作 1："acb" -> "bca"
示例 2：

输入：word1 = "a", word2 = "aa"
输出：false
解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
示例 3：

输入：word1 = "cabbba", word2 = "abbccc"
输出：true
解释：3 次操作从 word1 获得 word2 。
执行操作 1："cabbba" -> "caabbb"
执行操作 2："caabbb" -> "baaccc"
执行操作 2："baaccc" -> "abbccc"
示例 4：

输入：word1 = "cabbba", word2 = "aabbss"
输出：false
解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
思路：
若word2中的单词不在单词1中则直接返回false

*/
func closeStrings(word1 string, word2 string) bool {
	c1, c2 := make([]int, 26), make([]int, 26)

	for _, v := range word1 {
		c1[v-'a']++
	}
	for _, v := range word2 {
		c2[v-'a']++
	}
	for i := 0; i < 26; i++ {
		// c1中存在数据但是c2中不存在数据 或者 c1中的不存在数据c2中存在数据
		if c1[i] > 0 && c2[i] == 0 || c1[i] == 0 && c2[i] > 0 {
			return false
		}
	}
	sort.Ints(c1)
	sort.Ints(c2)

	return reflect.DeepEqual(c1, c2)
}
