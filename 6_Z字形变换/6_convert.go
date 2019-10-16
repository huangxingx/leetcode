package convert

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func convert(s string, numRows int) string {

	if len(s) < 2 || numRows < 2 {
		return s
	}

	step := 2*numRows - 2

	sList := make([]string, 0)

	for i := 0; i < len(s); i += step {
		sList = append(sList, s[i:Min(i+step, len(s))])
	}

	res := ""
	for _, value := range sList {
		res += string(value[0])
	}

	for i, j := 1, step-1; i <= j; i, j = i+1, j-1 {
		for _, value := range sList {
			if i < len(value) {
				res += string(value[i])
			}
			if j < len(value) && i != j {
				res += string(value[j])
			}
		}
	}

	return res
}
