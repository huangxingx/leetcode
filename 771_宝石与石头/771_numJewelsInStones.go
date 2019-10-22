package numJewelsInStones

import "strings"

func numJewelsInStones(J string, S string) int {
	ret := 0
	for _, j := range J {
		ret += strings.Count(S, string(j))
	}

	return ret
}
