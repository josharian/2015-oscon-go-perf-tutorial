// +build ignore

package main

import (
	"bytes"
	"fmt"
	"testing"
)

var p = bytes.Repeat([]byte{'a'}, 100)
var m = make(map[string]bool)

func BenchmarkMapKey1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m[string(p)] // HL
	}
}
func BenchmarkMapKey2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := string(p) // HL
		_ = m[s]       // HL
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
	run(BenchmarkMapKey1, "BenchmarkMapKey1")
	run(BenchmarkMapKey2, "BenchmarkMapKey2")
}
