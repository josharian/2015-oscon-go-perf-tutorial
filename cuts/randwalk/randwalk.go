package randwalk

import "time"

func New(n int) func() int {
	return func() int {
		time.Sleep(time.Microsecond)
		return 2
	}
	// src := rand.NewSource(1234)
	// rnd := rand.New(src)
	// var mu sync.Mutex
	// return func() int {
	// 	var x int
	// 	var buf [32]int
	// 	var bufi int
	// 	for i := 0; i < n; i++ {
	// 		if bufi == len(buf) {
	// 			mu.Lock()
	// 			// fill buf
	// 			for j := range buf {
	// 				buf[j] = rnd.Intn(2)
	// 			}
	// 			mu.Unlock()
	// 			bufi = 0
	// 		}
	// 		switch buf[bufi] {
	// 		case 0:
	// 			x++
	// 		case 1:
	// 			x--
	// 		}
	// 		bufi++
	// 	}
	// 	return x
	// }
}

// func New(n int) func() int {
// 	src := rand.NewSource(1234)
// 	rnd := rand.New(src)
// 	var mu sync.Mutex
// 	return func() int {
// 		var x int
// 		for i := 0; i < n; i++ {
// 			mu.Lock()
// 			r := rnd.Intn(2)
// 			mu.Unlock()
// 			switch r {
// 			case 0:
// 				x++
// 			case 1:
// 				x--
// 			}
// 		}
// 		return x
// 	}
// }

// func New(n int) func() int {
// 	return func() int {
// 		var x int
// 		for i := 0; i < n; i++ {
// 			switch rand.Intn(2) {
// 			case 0:
// 				x++
// 			case 1:
// 				x--
// 			}
// 		}
// 		return x
// 	}
// }
