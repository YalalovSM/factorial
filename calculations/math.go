package calculations

import (
	"math/big"
)

// FactorialNaive Calculate factorial by naive way
func FactorialNaive(n int) *big.Int {
	var res *big.Int = big.NewInt(1)
	for i := 2; i <= n; i++ {
		res = big.NewInt(0).Mul(res, big.NewInt(int64(i)))
	}

	return res
}

// FactorialTree Calculate factorial by tree calculation method
func FactorialTree(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}

	if n == 1 {
		return big.NewInt(1)
	}

	if n == 1 || n == 2 {
		return big.NewInt(int64(n))
	}

	return prodTree(2, n)
}

func prodTree(left, right int) *big.Int {
	if left > right {
		return big.NewInt(1)
	}

	if left == right {
		return big.NewInt(int64(left))
	}

	if right-left == 1 {
		return big.NewInt(0).Mul(big.NewInt(int64(left)), big.NewInt(int64(right)))
	}

	m := (left + right) / 2

	return big.NewInt(0).Mul(prodTree(left, m), prodTree(m+1, right))
}

// FactorialNaiveChannels calculate factorial by naive way but using goroutines
func FactorialNaiveChannels(n int) *big.Int {
	ch := make(chan *big.Int)

	go factorialNaiveChannels(n, big.NewInt(1), ch)

	return <-ch
}

func factorialNaiveChannels(n int, res *big.Int, ch chan *big.Int) {
	res = big.NewInt(0).Mul(res, big.NewInt(int64(n)))

	if n == 1 {
		ch <- res
		return
	}

	go factorialNaiveChannels(n-1, res, ch)
}
