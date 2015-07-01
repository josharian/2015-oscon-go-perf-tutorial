// +build ignore

package fib

func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 3
	}
	return Fib(n-1) + Fib(n-2)
}
