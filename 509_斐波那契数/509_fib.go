package fib

// 引入缓存
var sMap = make(map[int]int)

func fib(N int) int {

	if N <= 1 {
		return N
	}

	if r, ok := sMap[N]; ok {
		return r
	}
	ret := fib(N-1) + fib(N-2)
	sMap[N] = ret

	return ret
}
