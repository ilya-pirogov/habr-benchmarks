#!/usr/bin/env bash

set -e -o xtrace

cat /proc/meminfo

sudo bash -c "echo never > /sys/kernel/mm/transparent_hugepage/enabled"
perf stat -dddd ./habr-tests -timeout 10m $@

sudo bash -c "echo always > /sys/kernel/mm/transparent_hugepage/enabled"
perf stat -dddd ./habr-tests -timeout 10m $@
