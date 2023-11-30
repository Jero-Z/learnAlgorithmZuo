package hash

/**
205. 同构字符串

给定两个字符串 s 和 t ，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。
不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

输入：s = "egg", t = "add"
输出：true
输入：s = "foo", t = "bar"
输出：false
输入：s = "paper", t = "title"
输出：true
思路：
直接进行循环映射绑定
判断当前i位置的字符是否有映射关系，有则比较，如无则建立映射
思路：
建立两个hash表，保存相互之间的映射关系

*/
func isIsomorphic(s string, t string) bool {
	cnt1 := [255]uint8{}
	cnt2 := [255]uint8{}

	for i, v := range s {

		if cnt1[v] > 0 && cnt1[v] != t[i] || cnt2[t[i]] > 0 && cnt2[t[i]] != uint8(v) {
			return false
		}
		cnt1[v] = t[i]
		cnt2[t[i]] = uint8(v)
	}
	return true
}
