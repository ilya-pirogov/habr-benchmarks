package main

import (
	"flag"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var n uint64
	flag.Uint64Var(&n, "n", 1000000, "max N")
	flag.Parse()

	workers := runtime.NumCPU() / 2
	worker := func(arr []uint64, inners <-chan uint64) {
		for k1 := range inners {
			for k2 := k1; k2 <= n/k1; k2++ {
				var val uint64

				if k1 != k2 {
					val = k1 + k2
				} else {
					val = k1
				}

				atomic.AddUint64(&arr[k1*k2], val)
			}
		}
		wg.Done()
	}

	sqrtN := uint64(math.Sqrt(float64(n)))
    arr := make([]uint64, n+1)
    pool := make(chan uint64, sqrtN)

    for k1 := uint64(1); k1 <= sqrtN; k1++ {
        pool <- k1
    }
    close(pool)

    for i := 0; i < workers; i++ {
        wg.Add(1)
        go worker(arr, pool)
    }

    wg.Wait()

    println(arr[n])
}
