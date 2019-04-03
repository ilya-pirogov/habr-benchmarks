#!/usr/bin/env bash

set -e -o xtrace

sudo bash -c "echo never > /sys/kernel/mm/transparent_hugepage/enabled"
perf stat -dddd ./habr-benchmarks -timeout 10m 1e6 1e7 1e8 1e9 2> perf_thp_never.log

sudo bash -c "echo always > /sys/kernel/mm/transparent_hugepage/enabled"
perf stat -dddd ./habr-benchmarks -timeout 10m 1e6 1e7 1e8 1e9 2> perf_thp_always.log
