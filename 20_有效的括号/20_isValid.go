package isValid

func isValid(s string) bool {
	flagMap := map[string]string{")": "(", "}": "{", "]": "[",}
	stactList := make([]string, 0)
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		if si, ok := flagMap[string(s[i])]; len(stactList) > 0 && ok {
			if stactList[len(stactList)-1] != si {
				return false
			} else {
				// 如果匹配到一对的括号就出栈
				stactList = stactList[:len(stactList)-1]
			}
		} else {
			// 如果是正括号就入栈
			stactList = append(stactList, string(s[i]))
		}

	}

	return len(stactList) == 0

}
