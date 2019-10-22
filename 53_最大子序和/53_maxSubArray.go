package maxSubArray

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	ret := nums[0]

	maxNode := nums[0]

	for i := 1; i < len(nums); i++ {
		maxNode = max(nums[i]+maxNode, nums[i])
		ret = max(maxNode, ret)
	}

	return ret
}
