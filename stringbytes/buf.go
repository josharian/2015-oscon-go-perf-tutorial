// +build ignore

package main

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < 100; j++ {
			s += "a" // HL
		}
		_ = s
	}
}
func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for j := 0; j < 100; j++ {
			buf.WriteByte('a') // HL
		}
		_ = buf.String()
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
	run(BenchmarkConcat, "BenchmarkConcat")
	run(BenchmarkBuffer, "BenchmarkBuffer")
}
