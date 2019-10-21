package longestOnes

func longestOnes(A []int, K int) int {
	i, j, ret := 0, 0, 0

	for j < len(A) {
		if A[j] == 1 {
			j++
		} else {
			if K == 0 { // 没有容量了
				if A[i] == 0 { //排除尾部
					K++
				}
				i++
			} else {
				K--
				j++
			}
		}

		if ret < j-i {
			ret = j - i
		}
	}

	return ret
}
