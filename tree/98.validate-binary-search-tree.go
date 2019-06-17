/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	pre := -1 << 63
	tmp := root
	if tmp != nil || len(stack) != 0 {
		if tmp != nil {
			stack = append(stack, tmp)
			tmp = tmp.Left
		} else {
			tmp = stack[len(stack)-1]
			if pre >= tmp.Val {
				return false
			}
			stack = stack[:len(stack)-1]
			pre = tmp.Val
			tmp = tmp.Right
		}
	}
	return true
}

func isValidBST2(root *TreeNode) bool {
	var stack []*TreeNode
	// Go int类型最小值
	max := -1 << 63
	for nd := root; nd != nil || len(stack) != 0; {
		for nd != nil {
			stack = append(stack, nd)
			nd = nd.Left
		}
		if len(stack) != 0 {
			nd = stack[len(stack)-1]
			if nd.Val > max {
				max = nd.Val
			} else {
				return false
			}
			stack = stack[:len(stack)-1]
			nd = nd.Right
		}
	}
	return true
}
