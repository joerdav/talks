//go:build !quick

package arrays

func newHand() []float32 {
	return []float32{
		6,
		6.11,
		6.20,
		6.15,
		5.8,
	}
}

func averageFingerSize() float32 {
	hand := newHand()
	return (hand[0] +
		hand[1] +
		hand[2] +
		hand[3] +
		hand[4]) / float32(len(hand))
}

//goos: darwin
//goarch: amd64
//pkg: examples/arrays
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkSize-12        45860188                26.40 ns/op           24 B/op          1 allocs/op
