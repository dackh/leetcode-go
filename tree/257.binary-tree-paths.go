import "strconv"

/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root != nil {
		helper(root, "", &res)
	}
	return res
}

func helper(node *TreeNode, path string, res *[]string) {
	if node.Left == nil && node.Right == nil {
		*res = append(*res, path+strconv.Itoa(node.Val))
	}
	if node.Left != nil {
		helper(node.Left, path+strconv.Itoa(node.Val)+"->", res)
	}
	if node.Right != nil {
		helper(node.Right, path+strconv.Itoa(node.Val)+"->", res)
	}
}
