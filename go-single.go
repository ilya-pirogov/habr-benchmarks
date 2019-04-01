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

	for k1 := uint64(1); k1 <= uint64(math.Sqrt(float64(n))); k1++ {
		for k2 := k1; k2 <= n/k1; k2++ {
			var val uint64

			if k1 != k2 {
				val = k1 + k2
			} else {
				val = k1
			}

			arr[k1*k2] += val
		}
	}

	println(arr[n])
}
