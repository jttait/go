package popcount

import "testing"

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			PopCount(5)
		}
	}
}

func BenchmarkPopcount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			PopCount2(5)
		}
	}
}

func BenchmarkPopcount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			PopCount3(5)
		}
	}
}

func BenchmarkPopcount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			PopCount4(5)
		}
	}
}
