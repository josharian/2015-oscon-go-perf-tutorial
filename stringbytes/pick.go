// +build ignore

package main

import (
	"fmt"
	"testing"
)

var s string

func BenchmarkConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("foo-%d", i%3) // HL
	}
}
func BenchmarkPick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch i % 3 {
		case 0:
			s = "foo-0" // HL
		case 1:
			s = "foo-1" // HL
		case 2:
			s = "foo-2" // HL
		}
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
	run(BenchmarkConstruct, "BenchmarkConstruct")
	run(BenchmarkPick, "BenchmarkPick")
}
