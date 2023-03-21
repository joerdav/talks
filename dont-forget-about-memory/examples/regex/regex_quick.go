//go:build quick

package regex

import "regexp"

var (
	level2Heading = regexp.MustCompile("^-+$")
	level1Heading = regexp.MustCompile("^=+$")
)

func altHeadingLevel(nextLine string) int {
	if level2Heading.MatchString(nextLine) {
		return 2
	}
	if level1Heading.MatchString(nextLine) {
		return 1
	}
	return 0
}

// BenchmarkHeading2-12            14360788                75.66 ns/op            0 B/op          0 allocs/op
// BenchmarkHeading1-12            10675314               111.3 ns/op             0 B/op          0 allocs/op
// BenchmarkHeading0-12            16979590                69.95 ns/op            0 B/op          0 allocs/op
