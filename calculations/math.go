package calculations

import "math/big"

// NaiveFactorial Calculate factorial by naive way
func NaiveFactorial(n int) *big.Int {
	var res *big.Int = big.NewInt(1)
	for i := 2; i <= n; i++ {
		res = big.NewInt(0).Mul(res, big.NewInt(int64(i)))
	}

	return res
}
