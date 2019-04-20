package utils

import "math/big"

func fib(n int64, mems map[int64]*big.Int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}

	if n == 1 {
		return big.NewInt(1)
	}

	if mems[n] != nil {
		return mems[n]
	}

	r := big.NewInt(0)
	r.Add(fib(n-1, mems), fib(n-2, mems))
	mems[n] = r
	return mems[n]
}

func Fib(n int64) *big.Int {
	mems := map[int64]*big.Int{}
	return fib(n, mems)
}

func FibNumbers(n int64) []*big.Int {
	arr := []*big.Int{}
	i := int64(0)
	for {
		if i >= n {
			break
		}
		arr = append(arr, Fib(i))
		i++
	}

	return arr
}
