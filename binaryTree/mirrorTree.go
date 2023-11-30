package binaryTree

import "container/list"

/**
LCR 144. 翻转二叉树

给定一棵二叉树的根节点 root，请左右翻转这棵二叉树，并返回其根节点。
*/
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Front()).(*TreeNode)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		node.Left, node.Right = node.Right, node.Left
	}
	return root
}
func mirrorTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := mirrorTree(root.Left)
	root.Left = mirrorTree(root.Right)
	root.Right = temp
	return root
}
