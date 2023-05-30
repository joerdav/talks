//go:build !quick

package builder

func createAString() string {
	str := ""
	for i := 0; i < 100; i++ {
		str += "some string"
	}
	return str
}

// BenchmarkRandomGuids-12           105433             10166 ns/op           58872 B/op         99 allocs/op
