/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 *
 * https://leetcode.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Hard (27.68%)
 * Total Accepted:    159.5K
 * Total Submissions: 576.1K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * Given an array of non-negative integers, you are initially positioned at the
 * first index of the array.
 *
 * Each element in the array represents your maximum jump length at that
 * position.
 *
 * Your goal is to reach the last index in the minimum number of jumps.
 *
 * Example:
 *
 *
 * Input: [2,3,1,1,4]
 * Output: 2
 * Explanation: The minimum number of jumps to reach the last index is 2.
 * ⁠   Jump 1 step from index 0 to 1, then 3 steps to the last index.
 *
 * Note:
 *
 * You can assume that you can always reach the last index.
 *
 */
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var index, step int
	for index < len(nums) {
		if nums[index]+index >= len(nums)-1 {
			step++
			return step
		}
		max := 0
		nextIndex := index + 1
		for j := index + 1; j < len(nums) && j-index <= nums[index]; j++ {
			if max <= j+nums[j] {
				max = j + nums[j]
				nextIndex = j
			}
		}
		index = nextIndex
		step++
	}
	return step

}
