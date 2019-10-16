package plusOne

func plusOne(digits []int) []int {
	var isAddOne = true
	for i := len(digits) - 1; i >= 0; i-- {
		if isAddOne {
			digits[i] += 1
		} else {
			continue
		}
		if digits[i] >= 10 {
			isAddOne = true
		} else {
			isAddOne = false
		}

		digits[i] = digits[i] % 10

	}
	if isAddOne {
		return append([]int{1}, digits...)
	}

	return digits
}
