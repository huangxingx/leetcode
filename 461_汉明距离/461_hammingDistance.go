package hammingDistance

import (
	"strconv"
)

// 汉明距离
func hammingDistance(x int, y int) int {
	z := x ^ y

	zbin := strconv.FormatInt(int64(z), 2)
	sum := 0
	for _, value := range zbin {
		if "1" == string(value) {
			sum ++
		}
	}

	return sum

}
