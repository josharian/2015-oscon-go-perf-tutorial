// +build ignore

package main

import "testing"

func TestEncode(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{" ", "sp "},
		{"abc", " a  b  c "},
		{"ñ", ""},
	}

	e, err := load()
	if err != nil {
		t.Fatalf("failed to load ascii tables: %v", err)
	}

	for _, tt := range cases {
		got := e.encode(tt.in)
		if got != tt.want {
			t.Errorf("encode(%q)=%q want %q", tt.in, got, tt.want)
		}
	}
}

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		load()
	}
}

func BenchmarkEncode(b *testing.B) {
	b.ReportAllocs()
	const raw = "this is a moderately long string to encode"
	b.SetBytes(int64(len(raw)))
	e, err := load()
	if err != nil {
		b.Fatalf("failed to load ascii tables: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.encode(raw)
	}
}
