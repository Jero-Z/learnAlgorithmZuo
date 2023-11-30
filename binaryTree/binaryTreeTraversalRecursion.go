package binaryTree

import "container/list"

//levelOrder 层次遍历使用
func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	// 使用链表
	queue := list.New()
	queue.PushBack(root)
	var tempArr []int
	for queue.Len() > 0 {
		length := queue.Len() // 当前层的长度
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tempArr = append(tempArr, node.Val)
		}
		ret = append(ret, tempArr)
		tempArr = []int{}
	}
	return ret
}

//preOrderTraversal 前序遍历
func preOrderTraversal(root *TreeNode) []int {
	var values []int
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		values = append(values, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return values
}

//inOrderTraversal 中序遍历
func inOrderTraversal(root *TreeNode) []int {
	var traversal func(node *TreeNode)
	var values []int
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		values = append(values, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return values
}

//postOrderTraversal 后序遍历
func postOrderTraversal(root *TreeNode) []int {
	var traversal func(node *TreeNode)
	var values []int
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		traversal(node.Right)
		values = append(values, node.Val)
	}
	traversal(root)
	return values
}
func preOrderTraversal1(root *TreeNode) []int {
	var stack []*TreeNode
	var values []int

	node := root
	for node != nil || len(stack) > 0 {
		for node != nil { // 先根 后左
			values = append(values, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return values
}
