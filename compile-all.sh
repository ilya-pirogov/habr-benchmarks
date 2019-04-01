#!/usr/bin/env bash

set -e -o xtrace

go build ./habr-tests.go

g++ -O3 -march=native -o ./bin/cpp-single-gcc ./cpp-single.cpp
g++ -O3 -march=native -o ./bin/cpp-multi-gcc -pthread ./cpp-multi.cpp

ghc -O3 -o ./bin/haskell-single ./haskell-single.hs

go build -o ./bin/go-single -ldflags "-s -w" ./go-single.go
go build -o ./bin/go-multi-1 -ldflags "-s -w" ./go-multi-1.go
go build -o ./bin/go-multi-2 -ldflags "-s -w" ./go-multi-2.go

# have no idea how to solve dependencies without project
# rustc -o ./bin/rust-single ./rust-single.rs -- error[E0432]: unresolved import `clap`
# rustc -o ./bin/rust-multi ./rust-multi.rs
