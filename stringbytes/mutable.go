// +build ignore

package main

func set() {
	var b []byte
	b[0] = 0
	_ = b
}

func main() {
	set()
}
