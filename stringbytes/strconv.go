// +build ignore

package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(500) // HL
	}
}
func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", 500) // HL
	}
}

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
	run(BenchmarkStrconv, "BenchmarkStrconv")
	run(BenchmarkSprintf, "BenchmarkSprintf")
}
