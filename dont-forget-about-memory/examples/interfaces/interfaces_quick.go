//go:build quick

package interfaces

type calculator struct {
	add concreteAdder
}

func newCalculator() calculator {
	return calculator{
		add: concreteAdder{},
	}
}

func (c calculator) addSomeStuff(a, b int) int {
	return c.add.Add(a, b)
}

type concreteAdder struct{}

func (ca concreteAdder) Add(a, b int) int { return a + b }

// BenchmarkCalculator-12          310463655                3.881 ns/op           0 B/op          0 allocs/op

