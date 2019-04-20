package utils

func fib(n int, mems map[int]int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if mems[n] > 0 {
		return mems[n]
	}

	mems[n] = fib(n-1, mems) + fib(n-2, mems)
	return mems[n]
}

func Fib(n int) int {
	mems := map[int]int{}
	return fib(n, mems)
}

func FibNumbers(n int) []int {
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, Fib(i))
	}
	return arr
}
