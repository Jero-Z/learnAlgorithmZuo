package hash

/**
383. 赎金信

给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。
输入：ransomNote = "a", magazine = "b"
输出：false
输入：ransomNote = "aa", magazine = "ab"
输出：false
输入：ransomNote = "aa", magazine = "aab"
输出：true
思路：
若magazine 长度小于ransomNote 则一定不能构成
对ransomNote 进行遍历，每一次对magazine 进行消除，若所需的字符没有则表示不能组成
magazine 进行map[byte]int的构建
每次对其进行-- ，若小于1则不能成功消除，则抛出错误

*/
func canConstruct(ransomNote string, magazine string) bool {
	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}
	for _, ch := range ransomNote {
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return false
		}
	}
	return true

}
func canConstruct1(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	mcMap := make(map[int32]int, len(magazine))
	for _, i := range magazine {
		mcMap[i]++
	}
	for _, v := range ransomNote {
		if mcMap[v] < 1 {
			return false
		}
		mcMap[v]--
	}
	return true
}
