/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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
// func postorderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	postorder(root, &res)
// 	return res
// }

// func postorder(node *TreeNode, res *[]int) {
// 	if node != nil {
// 		postorder(node.Left, res)
// 		postorder(node.Right, res)
// 		*res = append(*res, node.Val)
// 	}
// }

//非递归
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := []*TreeNode{}
	tmp := root
	var rightNode *TreeNode = nil
	for tmp != nil || len(stack) != 0 {
		if tmp != nil {
			stack = append(stack, tmp)
			tmp = tmp.Left
		} else {
			node := stack[len(stack)-1]
			if node.Right != nil && node.Right != rightNode {
				tmp = node.Right
			} else {
				stack = stack[0 : len(stack)-1]
				res = append(res, node.Val)
				rightNode = node
			}
		}
	}
	return res
}
