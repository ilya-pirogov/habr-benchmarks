package main

import (
	"flag"
	"math"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var n uint64
	flag.Uint64Var(&n, "n", 1000000, "max N")
	flag.Parse()

	arr := make([]uint64, n+1)

	for k1 := uint64(1); k1 <= uint64(math.Sqrt(float64(n))); k1++ {
		wg.Add(1)
		go func(t1 uint64) {
			for k2 := t1; k2 <= n/t1; k2++ {
				var val uint64

				if t1 != k2 {
					val = t1 + k2
				} else {
					val = t1
				}

				atomic.AddUint64(&arr[t1*k2], val)
			}
			wg.Done()
		}(k1)
	}

	wg.Wait()

	println(arr[n])
}
