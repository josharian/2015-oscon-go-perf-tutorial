package flags

func repeat(b byte, n int) string {
	p := make([]byte, n)
	for i := 0; i < n; i++ {
		p[i] = b
	}
	return string(p)
}
