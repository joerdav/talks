//go:build !quick

package interfaces

func caclulateStuff(add adder, a, b int) int {
	return add.Add(a, b)
}

type adder interface {
	Add(a, b int) int
}

type concreteAdder struct{}

func (ca concreteAdder) Add(a, b int) int { return a + b }

// BenchmarkCalculator-12          191312350                6.231 ns/op           0 B/op          0 allocs/op
