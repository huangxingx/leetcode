package selfDividingNumbers

func selfDividingNumbers(left int, right int) []int {
	ret := make([]int, 0)
	for i := left; i < right+1; i++ {
		k := i
		flag := true
		for k > 0 {
			n := k % 10

			if n==0 || n != 0 && i%n != 0 {
				flag = false
				break
			}
			k /= 10
		}
		if flag {
			ret = append(ret, i)
		}
	}
	return ret
}
