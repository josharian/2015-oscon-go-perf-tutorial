package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type enc map[rune]string

func load() (enc, error) {
	f, err := os.Open("/usr/share/misc/ascii")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	e := make(enc)

	n := 0
	s := bufio.NewScanner(f)
	for s.Scan() && n < 128 {
		if len(s.Text()) == 0 {
			continue
		}
		off := 5
		for i := 0; i < 8 && n < 128; i++ {
			e[rune(n)] = s.Text()[off : off+3]
			off += 8
			n++
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}
	return e, nil
}

func (e enc) encode(s string) string {
	t := ""
	for _, r := range s {
		if r > 127 {
			continue
		}
		t += e[r]
	}
	return t
}

func main() {
	e, err := load()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.encode("hello\nprovincial world\a"))
}
