package lengthOfLongestSubstring

/*
遍历字符串通过map[字符串的byte]索引值
如果遍历的byte在map中存在,则通过获取到的重复byte的index+1 获取新数组,循环直到数组遍历完
 */

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {

	var tempMap map[byte]int
	var maxLength int
	var tempLength int
	sArray := []byte(s)
	var index int
	for {
		if index+1 > len(sArray) {
			break
		}

		tempMap = map[byte]int{}
		tempLength = 0

		for _, value := range sArray {
			index, ok := tempMap[value]
			if ok {
				break
			}
			tempMap[value] = index
			tempLength += 1
			if tempLength > maxLength {
				maxLength = tempLength
			}
		}
		sArray = sArray[index+1:]

	}

	return maxLength
}
