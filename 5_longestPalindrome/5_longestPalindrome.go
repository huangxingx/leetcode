package longestPalindrome

/*
中心扩展
考虑最长回文串为奇数和偶数两种情况
 */
func longestPalindrome(s string) string {

	//runtime.GOMAXPROCS(2)
	var start = 0
	var maxLen = 0
	sLen := len([]byte(s))
	if sLen < 2 {
		return s
	}
	result := make(chan [2]int, 2*sLen)

	for k := 0; k < sLen; k++ {
		go findMaxLength(k, k, []byte(s), result)   //假设最长为奇数，尝试尽可能向两边扩展
		go findMaxLength(k, k+1, []byte(s), result) //假设最长为偶数，尝试尽可能向两边扩展
	}

	for i := 0; i < 2*sLen; i++ {
		r := <-result
		//fmt.Println("result: ", strconv.Itoa(r[0]), "index: ", strconv.Itoa(r[1]))
		if maxLen < r[0] {
			maxLen = r[0]
			start = r[1]
			//fmt.Println("maxLen: ", maxLen)
		}
	}
	close(result)

	return string([]byte(s)[start : start+maxLen])
}

func findMaxLength(i, j int, bList []byte, retChan chan [2]int) {
	for i >= 0 && j < len(bList) && bList[i] == bList[j] {
		i -= 1
		j += 1
	}
	maxLength := j - i - 1
	startIndex := i + 1
	retList := [2]int{maxLength, startIndex}
	retChan <- retList

	return
}
