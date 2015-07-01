package fib

func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
