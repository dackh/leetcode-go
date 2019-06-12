/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 *
 * https://leetcode.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (40.60%)
 * Total Accepted:    304.9K
 * Total Submissions: 751.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, determine if it is height-balanced.
 *
 * For this problem, a height-balanced binary tree is defined as:
 *
 *
 * a binary tree in which the depth of the two subtrees of every node never
 * differ by more than 1.
 *
 *
 * Example 1:
 *
 * Given the following tree [3,9,20,null,null,15,7]:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * Return true.
 *
 * Example 2:
 *
 * Given the following tree [1,2,2,3,3,null,null,4,4]:
 *
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 *
 *
 * Return false.
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	_, ok := getDepth(root)
	return ok
}

func getDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, ok1 := getDepth(root.Left)
	rightDepth, ok2 := getDepth(root.Right)
	if !ok1 || !ok2 {
		return 0, false
	}
	if leftDepth-rightDepth > 1 || leftDepth-rightDepth < -1 {
		return 0, false
	}
	if leftDepth > rightDepth {
		return leftDepth + 1, true
	}
	return rightDepth + 1, true
}
