/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

package vscode

// @lc code=start
func removeElement(nums []int, val int) int {
	j := len(nums) - 1
	i := int(0)
	for i = 0; i <= j; i++ {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			i--
			j--
		}
	}
	return j + 1
}

// @lc code=end
