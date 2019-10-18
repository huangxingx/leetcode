package generate

// 动态规划
func generate(numRows int) [][]int {
	ret := make([][]int, 0)
	ret = append(ret, []int{1})

	if numRows == 0 {
		return [][]int{}
	} else if numRows == 1 {
		return ret
	}
	ret = append(ret, []int{1, 1})

	if numRows == 2 {
		return ret
	}

	for i := 2; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				// 第一个元素为1
				ret = append(ret, []int{1})

			} else if i == j {
				// 最后一个元素为1
				ret[i] = append(ret[i], 1)

			} else {
				ret[i] = append(ret[i], ret[i-1][j-1]+ret[i-1][j])

			}
		}
	}

	return ret
}
