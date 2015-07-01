package zipf

import "math/rand"

// func New(seed int64, n int) func() int {
// 	src := rand.NewSource(seed)
// 	r := rand.New(src)
// 	z := rand.NewZipf(r, 2, 1, uint64(n-1))
// 	var mu sync.Mutex
// 	return func() int {
// 		mu.Lock()
// 		defer mu.Unlock()
// 		return int(z.Uint64())
// 	}
// }

func New(seed int64, n int) func() int {
	src := rand.NewSource(seed)
	r := rand.New(src)
	z := rand.NewZipf(r, 2, 1, uint64(n-1))
	c := make(chan int, 32)
	go func() {
		for {
			c <- int(z.Uint64())
		}
	}()
	return func() int {
		return <-c
	}
}

// func New(seed int64, n int) func() int {
// 	const shards = 32
// 	var z [shards]*rand.Zipf
// 	for i := 0; i < shards; i++ {
// 		src := rand.NewSource(seed)
// 		r := rand.New(src)
// 		z[i] = rand.NewZipf(r, 2, 1, uint64(n-1))
// 	}
// 	var zmu [shards]sync.Mutex
// 	var x uint32
// 	return func() int {
// 		shard := int(atomic.AddUint32(&x, 1)) % shards
// 		zmu[shard].Lock()
// 		defer zmu[shard].Unlock()
// 		return int(z[shard].Uint64())
// 	}
// }
