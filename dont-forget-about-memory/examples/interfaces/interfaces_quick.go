//go:build quick

package interfaces

func caclulateStuff(add concreteAdder, a, b int) int {
	return add.Add(a, b)
}

type concreteAdder struct{}

func (ca concreteAdder) Add(a, b int) int { return a + b }

// BenchmarkCalculator-12          310463655                3.881 ns/op           0 B/op          0 allocs/op
