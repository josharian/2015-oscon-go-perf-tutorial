// +build ignore

package fib

import "testing"

func TestFib(t *testing.T) {
	// sequence from http://oeis.org/A000045
	seq := [...]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}
	for i, want := range seq {
		if got := Fib(i); got != want {
			t.Errorf("Fib(%d)=%d want %d", i, got, want)
		}
	}
}

var sink int

func benchmarkFib(b *testing.B, x int) {
	for i := 0; i < b.N; i++ {
		sink = Fib(x)
	}
}

func BenchmarkFib5(b *testing.B)  { benchmarkFib(b, 5) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib35(b *testing.B) { benchmarkFib(b, 35) }
