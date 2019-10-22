package game

func game(guess []int, answer []int) int {
	ret := 0
	for i := 0; i < len(answer); i++ {
		if guess[i] == answer[i] {
			ret++
		}

	}
	return ret
}
