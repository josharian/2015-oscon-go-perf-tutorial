package word

import (
	"bufio"
	"strings"
)

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
