package letterCombinations

var resultList = make([]string, 0)
var flagMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {

	if len(digits) != 0 {
		backtrack("", digits)
	}
	ret := resultList[:]
	resultList = make([]string, 0)
	return ret

}

func backtrack(re, nextString string) {
	if len(nextString) == 0 {
		resultList = append(resultList, re)
		return
	}

	currentNumber := string(nextString[0])
	currentString := flagMap[currentNumber]
	for _, value := range currentString {
		backtrack(re+string(value), nextString[1:])
	}

}
