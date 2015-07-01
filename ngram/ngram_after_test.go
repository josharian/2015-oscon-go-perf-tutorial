// +build ignore

package ngram

import (
	"strings"
	"testing"
)

func TestNgram(t *testing.T) {
	type W struct {
		s string
		n int
	}
	seq := [...]struct {
		add  string
		n    int
		want []W
	}{
		{add: "abc", n: 2, want: []W{W{s: "ab", n: 1}, W{s: "bc", n: 1}}},
		{add: "abc", n: 2, want: []W{W{s: "ab", n: 2}, W{s: "bc", n: 2}}},
		{add: "abc", n: 3, want: []W{W{s: "abc", n: 1}}},
	}

	c := New()
	for _, test := range seq {
		c.Add(test.add, test.n)
		for _, w := range test.want {
			if got := c.Get(w.s); got != w.n {
				t.Errorf("got %d want %d", got, w.n)
			}
		}
	}
}

func BenchmarkNGram(b *testing.B) {
	c := New()
	s := strings.Repeat("a man a plan a canal panama ", 15)
	b.SetBytes(int64(len(s)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Add(s, 3)
		}
	})
}
