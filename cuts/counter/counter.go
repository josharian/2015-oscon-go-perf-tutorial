package counter

import "sync"

const shards = 128

type Counter struct {
	mu [shards]sync.Mutex
	m  [shards]map[string]int
}

func NewCounter() *Counter {
	var c Counter
	for i := 0; i < shards; i++ {
		c.m[i] = make(map[string]int)
	}
	return &c
}

func (c *Counter) shard(s string) int {
	sum := 0
	for _, r := range s {
		sum += int(r)
	}
	return sum % shards
}

func (c *Counter) Add(s string) {
	x := c.shard(s)
	c.mu[x].Lock()
	defer c.mu[x].Unlock()
	c.m[x][s]++
}

func (c *Counter) Get(s string) int {
	x := c.shard(s)
	c.mu[x].Lock()
	defer c.mu[x].Unlock()
	return c.m[x][s]
}

func (c *Counter) Sum() int {
	for i := 0; i < shards; i++ {
		c.mu[i].Lock()
		defer c.mu[i].Unlock()
	}
	sum := 0
	for i := 0; i < shards; i++ {
		for _, n := range c.m[i] {
			sum += n
		}
	}
	return sum
}

// type Counter struct {
// 	mu sync.Mutex
// 	m  map[string]int
// }

// func NewCounter() *Counter {
// 	return &Counter{m: make(map[string]int)}
// }

// func (c *Counter) Add(s string) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.m[s]++
// }

// func (c *Counter) Get(s string) int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.m[s]
// }

// func (c *Counter) Sum() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	sum := 0
// 	for _, n := range c.m {
// 		sum += n
// 	}
// 	return sum
// }
