/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (46.74%)
 * Total Accepted:    193K
 * Total Submissions: 412.9K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 *
 * Example:
 *
 *
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 *
 */
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	helper(n, k, 1, cur, &res)
	return res

}

func helper(n, k, start int, cur []int, res *[][]int) {
	if k == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp) //二维切片不是引用传递
		return
	}
	for i := start; i <= n-k+1; i++ {
		cur = append(cur, i)
		helper(n, k-1, i+1, cur, res)
		cur = cur[:len(cur)-1]
	}
}
