/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for k, v := range nums {
		numMap[v] = k
	}
	value := int(0)
	for k, v := range nums {
		value = target - v
		if k2, ok := numMap[value]; ok && k2 != k {
			return []int{k, k2}
		}
	}
	return []int{}
}

// @lc code=end

