package main

import (
	"flag"
	"math"
)

func main() {
	var n uint64
	flag.Uint64Var(&n, "n", 1000000, "max N")
	flag.Parse()

	arr := make([]uint64, n+1)
	sqrtN := uint64(math.Sqrt(float64(n)))

	for k1 := uint64(1); k1 <= sqrtN; k1++ {
		for k2 := k1; k2 <= n/k1; k2++ {
			if k1 != k2 {
				arr[k1*k2] += k1 + k2
			} else {
				arr[k1*k2] += k1
			}
		}
	}

	println(arr[n])
}
