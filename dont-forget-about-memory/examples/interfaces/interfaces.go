//go:build !quick

package interfaces

type calculator struct {
	add adder
}

func newCalculator() calculator {
	return calculator{
		add: concreteAdder{},
	}
}

func (c calculator) addSomeStuff(a, b int) int {
	return c.add.Add(a, b)
}

type adder interface {
	Add(a, b int) int
}

type concreteAdder struct{}

func (ca concreteAdder) Add(a, b int) int { return a + b }

// BenchmarkCalculator-12          191312350                6.231 ns/op           0 B/op          0 allocs/op
