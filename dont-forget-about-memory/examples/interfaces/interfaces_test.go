package interfaces

import "testing"

func BenchmarkCalculator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculateStuff(concreteAdder{}, 1, 2)
	}
}
