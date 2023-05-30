//go:build vquick

package regex

func stringOnlyContains(input string, matcher rune) bool {
	if len(input) == 0 {
		return false
	}
	for i := range input {
		if []rune(input)[i] != matcher {
			return false
		}
	}
	return true
}

func altHeadingLevel(nextLine string) int {
	if stringOnlyContains(nextLine, '-') {
		return 2
	}
	if stringOnlyContains(nextLine, '=') {
		return 1
	}
	return 0
}

// BenchmarkHeading2-12            30728582                38.91 ns/op            0 B/op          0 allocs/op
// BenchmarkHeading1-12            22408417                53.66 ns/op            0 B/op          0 allocs/op
// BenchmarkHeading0-12            21599406                56.16 ns/op            0 B/op          0 allocs/op
