package zipf

import "testing"

func BenchmarkZipf(b *testing.B) {
	fn := New(1234, 10)
	// for i := 0; i < b.N; i++ {
	// 	_ = fn()
	// }
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fn()
		}
	})
}
