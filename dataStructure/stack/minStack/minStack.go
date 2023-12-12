package minStack

// 最小栈
// 测试链接 : https://leetcode.cn/problems/min-stack/

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.minStack) == 0 || val <= s.minStack[len(s.minStack)-1] {
		// 最小栈的位置调整
		s.minStack = append(s.minStack, val)
	}
}

func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1] // 对栈进行缩小

	if top == s.minStack[len(s.minStack)-1] {
		s.minStack = s.minStack[:len(s.minStack)-1] // 对最小栈进行缩小
	}

}

func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		return 0
	}
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.stack) == 0 {
		return 0
	}
	return s.minStack[len(s.minStack)-1]
}
