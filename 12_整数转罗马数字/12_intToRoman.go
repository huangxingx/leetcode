package intToRoman

// 整数转罗马数字
func intToRoman(num int) string {
	numList := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strList := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	ret := ""
	for num > 0 {
		for i := 0; i < len(numList); i++ {
			if num/numList[i] > 0 {
				for j := 0; j < num/numList[i]; j++ {
					ret += strList[i]
				}
				num %= numList[i]
			}
		}
	}

	return ret
}
