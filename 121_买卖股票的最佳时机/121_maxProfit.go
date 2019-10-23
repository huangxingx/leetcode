package maxProfit

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minV := prices[0]
	maxV := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minV{
			minV = prices[i]
		} else if prices[i]- minV > maxV{
			maxV = prices[i]- minV
		}
	}

	return maxV
}
