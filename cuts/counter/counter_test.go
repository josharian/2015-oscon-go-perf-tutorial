package counter

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	seq := [...]struct {
		s    string
		want int
	}{
		{s: "a", want: 1},
		{s: "a", want: 2},
		{s: "b", want: 1},
		{s: "", want: 1},
		{s: "a", want: 3},
	}

	c := NewCounter()
	for i, test := range seq {
		c.Add(test.s)
		n := c.Get(test.s)
		if n != test.want {
			t.Fatalf("added %q at step %d, got %d want %d", test.s, i, n, test.want)
		}
	}
}

var (
	once  sync.Once
	words []string
	err   error
)

func BenchmarkCounter(b *testing.B) {
	once.Do(func() {
		var buf []byte
		buf, err = ioutil.ReadFile("/usr/share/dict/connectives")
		if err != nil {
			return
		}
		for _, b := range bytes.Fields(buf) {
			words = append(words, string(b))
		}
	})

	if err != nil {
		b.Skipf("could not open dictionary: %v", err)
	}

	c := NewCounter()

	r := rand.New(rand.NewSource(1234))
	z := rand.NewZipf(r, 2, 1, uint64(len(words)-1))

	var seq [1024]int
	for i := 0; i < len(seq); i++ {
		// seq[i] = rand.Intn(len(words))
		seq[i] = int(z.Uint64())
	}
	_ = z

	b.ResetTimer()
	// for i := 0; i < b.N; i++ {
	// c.Add(words[b.N%len(words)])
	// c.Add(words[rand.Intn(len(words))])
	// }

	b.RunParallel(func(pb *testing.PB) {
		var i int
		for pb.Next() {
			// c.Add(words[rand.Intn(len(words))])
			// c.Add(words[int(z.Uint64())])
			c.Add(words[seq[i]])
			i = (i + 1) % 1024
		}
	})

	b.StopTimer()
	if got := c.Sum(); got != b.N {
		b.Errorf("Sum=%d want %d", got, b.N)
	}
}

// 	s := rand.NewSource(1234)
// 	r := rand.New(s)
// 	z := rand.NewZipf(r, 2, 1, uint64(len(n)-1))
