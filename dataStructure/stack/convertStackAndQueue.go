package stack

// 用栈实现队列
type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{stack: []int{}}
}

func (q *MyQueue) Push(x int) {
	q.stack = append(q.stack, x)
}

func (q *MyQueue) Pop() int {
	if len(q.stack) == 0 {
		return 0
	}
	first := q.stack[0]
	q.stack = q.stack[0:]
	return first
}

//Peek 查看队列头部的内容
func (q *MyQueue) Peek() int {
	if len(q.stack) == 0 {
		return 0
	}
	return q.stack[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.stack) == 0
}
