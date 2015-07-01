// +build ignore

package word

import (
	"reflect"
	"strings"
	"testing"
)

func TestLengths(t *testing.T) {
	cases := [...]struct {
		s    string
		want map[int]int
	}{
		{s: "", want: map[int]int{}},
		{s: "abc", want: map[int]int{3: 1}},
		{s: "abc ", want: map[int]int{3: 1}},
		{s: "abc def", want: map[int]int{3: 2}},
		{s: "a bc def", want: map[int]int{1: 1, 2: 1, 3: 1}},
	}

	for _, test := range cases {
		if got := Lengths(test.s); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Lengths(%q)=%v want %v", test.s, got, test.want)
		}
	}
}

var sink map[int]int

func BenchmarkWordLengths(b *testing.B) {
	const repeats = 50
	s := strings.Repeat("1 22 333 4444", repeats)
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = Lengths(s)
		// b.StopTimer()
		// b.Logf("huh? %v", sink)
		// if !reflect.DeepEqual(sink, map[int]int{1: 50, 2: 50, 3: 50, 4: 50}) {
		// 	b.Error("oops")
		// }
		// b.StartTimer()
	}
}
