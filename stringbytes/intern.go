// +build ignore

package main

import (
	"bytes"
	"fmt"
	"testing"
)

var interned = make(map[string]string)

func intern(b []byte) string {
	s, ok := interned[string(b)] // does not allocate! // HL
	if ok {
		return s
	}
	s = string(b)
	interned[s] = s
	return s
}

var p = bytes.Repeat([]byte{'a'}, 100)

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(p) // HL
	}
}
func BenchmarkIntern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = intern(p) // HL
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
	run(BenchmarkConvert, "BenchmarkConvert")
	run(BenchmarkIntern, "BenchmarkIntern")
}
