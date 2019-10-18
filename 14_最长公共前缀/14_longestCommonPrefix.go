package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}
	flag := true
	ret := make([]byte, 0)
	for i := 0; i < len(strs[0]); i++ {
		for _, str := range strs[1:] {
			if i >= len(str) {
				flag = false
				break
			}

			if strs[0][i] != str[i] {
				flag = false
				break
			}
		}
		if !flag {
			break
		}

		ret = append(ret, strs[0][i])

	}

	return string(ret)
}
