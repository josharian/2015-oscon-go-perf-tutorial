// +build ignore

package ngram

import (
	"strings"
	"sync"
	"unicode"
)

type Count struct {
	mu sync.Mutex
	m  map[string]int
}

func New() *Count {
	return &Count{m: make(map[string]int)}
}

func (c *Count) Add(s string, n int) {
	x := make(map[string]int)
	s = strings.Map(asciilower, s)
	words := strings.Fields(s)
	for _, word := range words {
		for i := 0; i < len(word)-n+1; i++ {
			x[s[i:i+n]]++
		}
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range x {
		c.m[k] += v
	}
}

func (c *Count) Get(s string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.m[s]
}

func asciilower(r rune) rune {
	if r > 127 {
		return -1
	}
	if unicode.IsLetter(r) {
		return unicode.ToLower(r)
	}
	if unicode.IsSpace(r) {
		return r
	}
	return -1
}
