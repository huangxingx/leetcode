package myAtoi

import (
	"regexp"
	"strconv"
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	b := []byte(str)
	pat := `^[+|-]?\d+`
	reg := regexp.MustCompile(pat)
	newStr := reg.Find(b)
	number, _ := strconv.Atoi(string(newStr))

	if number > math.MaxInt32 {
		return math.MaxInt32
	} else if number < math.MinInt32 {
		return math.MinInt32
	}

	return number
}
