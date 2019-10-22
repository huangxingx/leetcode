package rob

/*
动态规划
最大金额: n[i] = n[i-2] + n[i]
 */

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 打家劫舍
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	moneys := make([]int, len(nums))
	moneys[0] = nums[0]
	moneys[1] = nums[1]
	maxMoney := max(moneys[0], moneys[1])

	for i := 2; i < len(nums); i++ {

		temp := moneys[i-2]
		if i-3 >= 0 {
			temp = max(moneys[i-2], moneys[i-3])
		}
		moneys[i] = temp + nums[i]

		maxMoney = max(maxMoney, moneys[i])
	}

	return maxMoney
}
