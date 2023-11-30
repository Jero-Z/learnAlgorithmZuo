package string

import (
	"fmt"
	"strings"
)

/*
151. 反转字符串中的单词

给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。



示例 1：

输入：s = "the sky is blue"
输出："blue is sky the"
示例 2：

输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。
示例 3：

输入：s = "a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。

思路：
遍历字符串，
使用一个临时变量和一个栈
遇到空格进行字符串的截断
临时字符串不为空的话则入栈
最后出栈，最后一个单元后不追加空格
返回拼接字符串
*/

func reverseWords(s string) string {
	t := strings.Fields(s)          //使用该函数可切割单个/多个空格，提取出单词
	for i := 0; i < len(t)/2; i++ { //遍历数组的前半段直接交换即可
		t[i], t[len(t)-1-i] = t[len(t)-1-i], t[i]
	}
	return strings.Join(t, " ") //重新使用空格连接多个单词
}

func reverseWords1(s string) string {
	var stack []string
	temp := ""
	for _, i := range s {
		fmt.Println(i)
		if i == ' ' {
			if len(temp) != 0 {
				stack = append(stack, temp)
				temp = ""
			}
			continue
		}

		temp += string(i)
	}
	if temp != "" {
		stack = append(stack, temp)
	}
	if len(stack) == 1 {
		return stack[0]
	}
	res := ""
	sc := len(stack) - 1
	for i := sc; i >= 0; i-- {
		if i > 0 {
			res += string(stack[i]) + " "
			continue
		}
		res += stack[i]
	}
	return res
}
