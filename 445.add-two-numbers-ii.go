/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 *
 * https://leetcode.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (49.59%)
 * Total Accepted:    84.5K
 * Total Submissions: 170.4K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The most significant digit comes first and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Follow up:
 * What if you cannot modify the input lists? In other words, reversing the
 * lists is not allowed.
 *
 *
 *
 * Example:
 *
 * Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 8 -> 0 -> 7
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	head := &ListNode{
		Val: -1,
	}
	count := 0
	for len(stack1) != 0 || len(stack2) != 0 || count != 0 {
		sum := 0
		if len(stack1) != 0 {
			sum += stack1[len(stack1)-1]
			stack1 = stack1[0 : len(stack1)-1]
		}
		if len(stack2) != 0 {
			sum += stack2[len(stack2)-1]
			stack2 = stack2[0 : len(stack2)-1]
		}
		sum += count
		tmp := &ListNode{
			Val:  sum % 10,
			Next: head.Next,
		}
		head.Next = tmp
		count = sum / 10
	}
	return head.Next
}
