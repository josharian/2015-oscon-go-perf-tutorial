package flags

func clear(b []byte) {
	for i := range b {
		b[i] = 0
	}
}
