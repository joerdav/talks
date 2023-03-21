//go:build quick

package builder

import (
	"strings"
)

func createAString() string {
	var build strings.Builder
	for i := 0; i < 100; i++ {
		build.WriteString("some string")
	}
	return build.String()
}

// BenchmarkRandomGuids-12           976308              1225 ns/op            3312 B/op          8 allocs/op

