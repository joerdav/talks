package interfaces

import "testing"

func BenchmarkCalculator(b *testing.B) {
	calc := newCalculator()
	for i := 0; i < b.N; i++ {
		calc.addSomeStuff(1, 2)
	}
}
