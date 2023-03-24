package util

import (
	"golang.org/x/exp/constraints"
)

func Min[N constraints.Integer | constraints.Float](a, b N) N {
	if a < b {
		return a
	}
	return b
}

func Max[N constraints.Integer | constraints.Float](a, b N) N {
	if a > b {
		return a
	}
	return b
}

func Factorial[N constraints.Integer](n N) N {
	var fn N = 1
	for n > 1 {
		fn *= n
		n--
	}
	return fn
}

func Binomial[N constraints.Integer](n, k N) N {
	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}
