// +build ignore

package word

import (
	"bufio"
	"strings"
	"unicode"
)

func Lengths(s string) map[int]int {
	m := make(map[int]int)
	for {
		i := strings.IndexFunc(s, unicode.IsSpace)
		if i == -1 {
			if len(s) > 0 {
				m[len(s)]++
			}
			break
		}
		m[i]++
		s = s[i:]
		s = strings.TrimSpace(s)
	}
	return m
}

func Lengths(s string) map[int]int {
	m := make(map[int]int)
	var n int
	for _, r := range s {
		if !unicode.IsSpace(r) {
			n++
			continue
		}
		if n > 0 {
			m[n]++
		}
		n = 0
	}
	if n > 0 {
		m[n]++
	}
	return m
}

func Lengths(s string) map[int]int {
	m := make(map[int]int)
	ss := strings.Fields(s)
	for _, x := range ss {
		m[len(x)]++
	}
	return m
}

func Lengths(s string) map[int]int {
	m := make(map[int]int)
	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Bytes()
		m[len(word)]++
	}
	return m
}
