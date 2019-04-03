## Benchmark collection form https://habr.com/en/post/445134/

### Requirements

- Docker for building all binaries

### Build

```
make all
```

### Usage

```
habr-benchmarks [options] 1e6 1e7 1e8 ...
  -filter string
    	execute only tests with specific name
  -perf
    	use perf tool for each test (default true)
  -timeout duration
    	timout for each test (default 10m0s)

```

