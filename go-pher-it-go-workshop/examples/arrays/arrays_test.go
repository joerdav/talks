package arrays

import "testing"

func BenchmarkSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = averageFingerSize()
	}
}
