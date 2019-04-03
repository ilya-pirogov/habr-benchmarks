DOCKER=docker run --rm -v /tmp:/tmp -e HOME=/tmp/habr-benchmarks -e CFLAGS='$(CFLAGS)' -v $(CURDIR):/mnt -u $$(id -u):$$(id -g) --workdir /mnt
CFLAGS=-O3 -march=native -pthread

all: habr-benchmarks bin/cpp-single bin/cpp-multi bin/go-single bin/go-multi-1 bin/go-multi-2 bin/haskell-single bin/rust-single bin/rust-multi bin/zig-single

habr-benchmarks:
	$(DOCKER) golang:1.12 go build -ldflags "-s -w" $< -o $@

bin/%: cpp/%.cpp
	$(DOCKER) gcc:8.3 g++ $< -o $@ $(CFLAGS)

bin/%: go/%.go
	$(DOCKER) golang:1.12 go build -ldflags "-s -w" -o $@ $<

bin/%: haskell/%.hs
	$(DOCKER) haskell:8.6 ghc -O3 $< -o $@

bin/%: rust/%.rs
	$(DOCKER) rust:1.33 bash -c 'cd rust && cargo build --release && cp ./target/release/$* ../$@'

bin/%: zig/%.zig
	docker build -t zig:0.3 zig
	$(DOCKER) zig:0.3 zig --release-fast --single-threaded --strip --name $@ build-exe $<

clean:
	rm -f habr-benchmarks bin/*
	rm -rf rust/target
	rm -rf /tmp/habr-benchmarks

clean-images:
	docker image prune
	docker rmi zig

.PHONY: clean clean-images
