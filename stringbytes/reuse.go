// +build ignore

package main

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

var pool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}
var p = bytes.Repeat([]byte{'a'}, 100)

func BenchmarkReuse(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := pool.Get().(*bytes.Buffer) // HL
			buf.Write(p)
			_ = buf.String()
			buf.Reset()   // HL
			pool.Put(buf) // HL
		}
	})
}
func BenchmarkNoReuse(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var buf bytes.Buffer // HL
			buf.Write(p)
			_ = buf.String()
		}
	})
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
	run(BenchmarkReuse, "BenchmarkReuse")
	run(BenchmarkNoReuse, "BenchmarkNoReuse")
}
