package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/dustin/go-humanize"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

var implementations = []Implementation{
	{"C++: gcc -O3 -march-native; single thread", "./bin/cpp-single %d"},
	{"C++: gcc -O3 -march-native; 10 threads", "./bin/cpp-multi %d 10"},
	{"Haskell: ghc -O3; single thread", "./bin/haskell-single %d"},
	{"Go: single thread", "./bin/go-single -n %d"},
	{"Go: sqrt(n) parallel goroutines", "./bin/go-multi-1 -n %d"},
	{"Go: NumCPU/2 goroutines + chan", "./bin/go-multi-2 -n %d"},
	{"Rust: release; single thread", "./bin/rust-single %d"},
	{"Rust: release; multi threads", "./bin/rust-multi %d"},
	{"Zig: zig --release-fast; single thread", "./bin/zig-single %d"},
}
var maxNumbers []uint64
var timeout time.Duration
var usePerf bool
var filter string

type Implementation struct {
	Name string
	Cmd  string
}
type Result table.Row
type ResultList []Result

func (r *Result) Test(n uint64, tpl string) {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	defer func(s time.Time) {
		if ctx.Err() != nil {
			r.Log(ctx.Err())
		} else if err != nil {
			r.Log(err)
		} else {
			r.Log(s)
		}
	}(time.Now())

	cmd := fmt.Sprintf(tpl, n)

	if usePerf {
		cmd = "perf stat -dddd " + cmd
	}

	args := strings.Fields(cmd)
	ex := exec.CommandContext(ctx, args[0], args[1:]...)
	ex.Stderr = os.Stderr
	err = ex.Run()

	if err != nil && ctx.Err() != context.DeadlineExceeded {
		fmt.Printf("\n%s: %s\n", cmd, err.Error())
	}
	print(".")
}

func (r *Result) Log(res interface{}) {
	var newResult Result
	switch v := res.(type) {
	case error:
		if v == context.DeadlineExceeded {
			newResult = append(*r, "timeout")
		} else {
			newResult = append(*r, "error")
		}
	case time.Time:
		newResult = append(*r, time.Now().Sub(v))
	}
	*r = newResult
}

func (res ResultList) Len() int      { return len(res) }
func (res ResultList) Swap(i, j int) { res[i], res[j] = res[j], res[i] }
func (res ResultList) Less(i, j int) bool {
	l, lok := res[i][len(maxNumbers)].(time.Duration)
	r, rok := res[j][len(maxNumbers)].(time.Duration)

	if !lok {
		return false
	}

	if !rok {
		return true
	}

	return l < r
}

func printTable(results []Result) {
	tbl := table.NewWriter()
	tbl.SetAutoIndex(true)
	tbl.SetOutputMirror(os.Stdout)
	header := table.Row{"Implementation"}
	for _, n := range maxNumbers {
		header = append(header, humanize.Comma(int64(n)))
	}
	tbl.AppendHeader(header)

	for _, impl := range results {
		row := table.Row{}
		for _, val := range impl {
			switch r := val.(type) {
			case time.Duration:
				row = append(row, r.Round(time.Millisecond))
			default:
				row = append(row, r)
			}
		}
		tbl.AppendRow(row)
	}

	tbl.Render()
}

func main() {
	flag.DurationVar(&timeout, "timeout", 10*time.Minute, "timout for each test")
	flag.BoolVar(&usePerf, "perf", true, "use perf tool for each test")
	flag.StringVar(&filter, "filter", "", "execute only tests with specific name")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Printf("usage: %s 1e6 1e7 1e8 ...\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	for _, v := range flag.Args() {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("unable to parse %s arg: %s\n", v, err)
			os.Exit(2)
		}
		maxNumbers = append(maxNumbers, uint64(n))
	}

	var filtered []Implementation
	if filter == "" {
		filtered = implementations
	}

	filtered = make([]Implementation, 0)
	for _, impl := range implementations {
		if strings.Contains(impl.Name, filter) {
			filtered = append(filtered, impl)
		}
	}
	results := make(ResultList, len(filtered))

	for i, impl := range filtered {
		res := Result{impl.Name}
		for _, n := range maxNumbers {
			res.Test(n, impl.Cmd)
		}
		results[i] = res
	}

	sort.Sort(results)
	println()
	printTable(results)
}
