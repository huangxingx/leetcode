package isPalindrome

// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
func isPalindrome(x int) bool {
	// 负数肯定不是回文数
	if x < 0 {
		return false
	}
	var result int
	newX := x
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	if result != newX {
		return false
	}

	return true
}

func isPalindrome2(x int) bool {
	// 负数肯定不是回文数
	if x < 0 {
		return false
	}
	resultList := make([]int, 0)
	for x != 0 {
		resultList = append(resultList, x%10)
		x /= 10
	}
	rLen := len(resultList)
	for i := 0; i < rLen/2; i++ {
		if resultList[i] != resultList[rLen-1-i] {
			return false
		}
	}

	return true

}
