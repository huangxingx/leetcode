package isHappy

import (
	"math"
)

func isHappy(n int) bool {
	count := 0
	for {
		ret := 0.0
		for n > 0 {
			x := n % 10
			ret += math.Pow(float64(x), float64(2))
			n /= 10
		}

		if ret == float64(1) {
			return true
		}
		n = int(ret)

		count++
		if count > 1000 {
			return false
		}
	}

}
