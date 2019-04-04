/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (28.10%)
 * Total Accepted:    846.9K
 * Total Submissions: 3M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 *
 *
 *
 *
 */
func lengthOfLongestSubstring(s string) int {
	var arr = make([]byte, 0)
	var i, j, res = 0, 0, 0
	for i < len(s) && j < len(s) {
		if isExist(arr, s[j]) {
			arr = arr[1:]
			i++
		} else {
			arr = append(arr, s[j])
			j++
			if res < len(arr) {
				res = len(arr)
			}
		}
	}
	return res
}

func isExist(arr []byte, key byte) bool {
	for _, value := range arr {
		if value == key {
			return true
		}
	}
	return false
}
