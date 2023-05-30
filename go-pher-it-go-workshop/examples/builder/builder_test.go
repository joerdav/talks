package builder

import "testing"

func BenchmarkRandomGuids(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = createAString()
	}
}
