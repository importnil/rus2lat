package rus2lat

import (
	"testing"
)

func BenchmarkRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Raw("Привет, я Вася!")
	}
}

func BenchmarkURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		URL("Привет, я Вася!")
	}
}
