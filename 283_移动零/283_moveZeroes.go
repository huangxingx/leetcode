package moveZeroes

// 移动零
func moveZeroes(nums []int) {
	zeroeCount := 0
	for i, value := range nums {
		if value == 0 {
			zeroeCount ++
			continue
		}
		if zeroeCount > 0 {
			nums[i], nums[i-zeroeCount] = nums[i-zeroeCount], nums[i]
		}
	}
}
