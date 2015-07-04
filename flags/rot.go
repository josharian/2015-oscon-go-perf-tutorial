package flags

func rot13(i uint64) uint64 {
	return i<<13 | i>>(64-13)
}

func rot13slow(i uint64) uint64 {
	i <<= 13
	i >>= 64 - 13
	return i
}
