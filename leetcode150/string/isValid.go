package string

/**
20. 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
思路：
使用队列
如果是左括号入栈
如果是碰到右括号则进行左闭右开的进行[:i-1]匹配
*/
func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 { // 字符的长度为奇数则肯定不能进行匹配
		return false
	}

	hash := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			//进行入栈
			stack = append(stack, s[i])

		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
			//如果栈中存在数据，并且栈中的数据为最后一个左括号 开始进行反括号的匹配
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
