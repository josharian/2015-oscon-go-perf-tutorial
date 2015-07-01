// +build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strings"
)

var ascii = `
|000 nul|001 soh|002 stx|003 etx|004 eot|005 enq|006 ack|007 bel|
|010 bs |011 ht |012 nl |013 vt |014 np |015 cr |016 so |017 si |
|020 dle|021 dc1|022 dc2|023 dc3|024 dc4|025 nak|026 syn|027 etb|
|030 can|031 em |032 sub|033 esc|034 fs |035 gs |036 rs |037 us |
|040 sp |041  ! |042  " |043  # |044  $ |045  % |046  & |047  ' |
|050  ( |051  ) |052  * |053  + |054  , |055  - |056  . |057  / |
|060  0 |061  1 |062  2 |063  3 |064  4 |065  5 |066  6 |067  7 |
|070  8 |071  9 |072  : |073  ; |074  < |075  = |076  > |077  ? |
|100  @ |101  A |102  B |103  C |104  D |105  E |106  F |107  G |
|110  H |111  I |112  J |113  K |114  L |115  M |116  N |117  O |
|120  P |121  Q |122  R |123  S |124  T |125  U |126  V |127  W |
|130  X |131  Y |132  Z |133  [ |134  \ |135  ] |136  ^ |137  _ |
|140  ` + "`" + ` |141  a |142  b |143  c |144  d |145  e |146  f |147  g |
|150  h |151  i |152  j |153  k |154  l |155  m |156  n |157  o |
|160  p |161  q |162  r |163  s |164  t |165  u |166  v |167  w |
|170  x |171  y |172  z |173  { |174  | |175  } |176  ~ |177 del|
`

type enc [128][]byte

func load() (enc, error) {
	f := strings.NewReader(ascii)

	var e enc

	n := 0
	s := bufio.NewScanner(f)
	for s.Scan() && n < 128 {
		b := s.Bytes()
		if len(b) == 0 {
			continue
		}
		bb := make([]byte, len(b))
		copy(bb, b)
		off := 5
		for i := 0; i < 8 && n < 128; i++ {
			e[n] = bb[off : off+3]
			off += 8
			n++
		}
	}

	if err := s.Err(); err != nil {
		return e, err
	}
	return e, nil
}

func (e enc) encode(s string) string {
	var b bytes.Buffer
	for _, r := range s {
		if r > 127 {
			continue
		}
		b.Write(e[r])
	}
	return b.String()
}

func main() {
	e, err := load()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.encode("hello\nprovincial world\a"))
}
