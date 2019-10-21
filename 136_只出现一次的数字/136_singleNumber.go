package singleNumber

// 位操作
func singleNumber(nums []int) int {
	ret := 0
	for _, value := range nums {
		ret ^= value
	}
	return ret
}

/*
复杂度分析

时间复杂度： O(n) 。我们只需要将 nums 中的元素遍历一遍，所以时间复杂度就是 nums 中的元素个数。
空间复杂度：O(1)。

 */
