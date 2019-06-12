/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归
// func inorderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	inorder(root, &res)
// 	return res
// }

// func inorder(node *TreeNode, res *[]int) {
// 	if node != nil {
// 		inorder(node.Left, res)
// 		*res = append(*res, node.Val)
// 		inorder(node.Right, res)
// 	}
// }

//非递归
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := []*TreeNode{}
	tmp := root
	for tmp != nil || len(stack) != 0 {
		if tmp != nil {
			stack = append(stack, tmp)
			tmp = tmp.Left
		} else {
			tmp = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			res = append(res, tmp.Val)
			tmp = tmp.Right
		}
	}
	return res
}
