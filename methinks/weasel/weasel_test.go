package weasel

import "testing"

func BenchmarkWeasel1(b *testing.B) {
	for b.Loop() {
		Weasel1()
	}
}

func BenchmarkWeasel2(b *testing.B) {
	for b.Loop() {
		Weasel2()
	}
}
