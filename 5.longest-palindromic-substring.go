/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (26.84%)
 * Total Accepted:    505.1K
 * Total Submissions: 1.9M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: "cbbd"
 * Output: "bb"
 *
 *
 */
package main

func main() {
	longestPalindrome("cbbd")
}

var start, end int

func longestPalindrome(s string) string {
	for i := 0; i < len(s); i++ {
		helper(s, i, i)
		helper(s, i, i+1)
	}
	return s[start:end]
}

func helper(s string, lo, ro int) {
	for lo >= 0 && ro < len(s) && s[lo] == s[ro] {
		lo--
		ro++
	}
	if end-start < ro-lo-1 {
		start = lo + 1
		end = ro
	}
}
