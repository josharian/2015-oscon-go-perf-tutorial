// +build ignore

package main

import (
	"fmt"
	"testing"
)

var sink []int

const size = 100

func BenchmarkDelayedAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int // HL
		for i := 0; i < size; i++ {
			s = append(s, i)
		}
		sink = s
	}
}
func BenchmarkOneAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 100) // HL
		for i := 0; i < size; i++ {
			s = append(s, i)
		}
		sink = s
	}
}
func _() {} // OMIT

func run(fn func(*testing.B), name string) {
	r := testing.Benchmark(fn)
	fmt.Println(name)

	nsop := r.NsPerOp()
	ns := fmt.Sprintf("%d ns/op", nsop)
	if r.N > 0 && nsop < 100 {
		if nsop < 10 {
			ns = fmt.Sprintf("%.2f ns/op", float64(r.T.Nanoseconds())/float64(r.N))
		} else {
			ns = fmt.Sprintf("%.1f ns/op", float64(r.T.Nanoseconds())/float64(r.N))
		}
	}

	fmt.Printf("\t%s\n", ns)
	fmt.Printf("\t%d B/op\n", r.AllocedBytesPerOp())
	fmt.Printf("\t%d allocs/op\n", r.AllocsPerOp())
}

func main() {
	run(BenchmarkDelayedAlloc, "BenchmarkDelayedAlloc")
	run(BenchmarkOneAlloc, "BenchmarkOneAlloc")
}
