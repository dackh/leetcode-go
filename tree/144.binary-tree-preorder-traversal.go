/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := []*TreeNode{}
	tmp := root
	for tmp != nil || len(stack) != 0 {
		if tmp != nil {
			stack = append(stack, tmp)
			res = append(res, tmp.Val)
			tmp = tmp.Left
		} else {
			tmp = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			tmp = tmp.Right
		}
	}
	return res
}
