package regex

import "testing"

func BenchmarkHeading2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		altHeadingLevel("---")
	}
}

func BenchmarkHeading1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		altHeadingLevel("===")
	}
}

func BenchmarkHeading0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		altHeadingLevel("not a heading")
	}
}
