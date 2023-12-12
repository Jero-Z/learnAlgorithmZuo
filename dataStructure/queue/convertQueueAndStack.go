package queue

// 225. 用队列实现栈

type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{q: []int{}}
}

func (s *MyStack) Push(x int) {
	s.q = append(s.q, x)

}

func (s *MyStack) Pop() int {
	if len(s.q) == 0 {
		return 0
	}
	top := s.q[len(s.q)-1]
	s.q = s.q[0 : len(s.q)-1]
	return top
}

func (s *MyStack) Top() int {
	if len(s.q) == 0 {
		return 0
	}
	return s.q[len(s.q)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.q) == 0
}
