//go:build !(quick || vquick)

package regex

import "regexp"

func altHeadingLevel(nextLine string) int {
	if regexp.MustCompile("^-+$").MatchString(nextLine) {
		return 2
	}
	if regexp.MustCompile("^=+$").MatchString(nextLine) {
		return 1
	}
	return 0
}

// BenchmarkHeading2-12              546982              2173 ns/op            2377 B/op         35 allocs/op
// BenchmarkHeading1-12              275372              4334 ns/op            4754 B/op         70 allocs/op
// BenchmarkHeading0-12              275092              4263 ns/op            4754 B/op         70 allocs/op
