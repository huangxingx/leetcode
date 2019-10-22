package flipAndInvertImage

func reverse(x int) int {
	if x == 0 {
		return 1
	}
	return 0
}

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if j < len(A[i])-j {
				A[i][j], A[i][len(A[i])-1-j] = reverse(A[i][len(A[i])-1-j]), reverse(A[i][j])
				//temp := A[i][len(A[i])-1-j]
				//A[i][len(A[i])-1-j] = reverse(A[i][j])
				//A[i][j] = reverse(temp)

			}
		}
	}
	return A
}
