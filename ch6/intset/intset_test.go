package intset

import (
	"testing"
	"math/rand"
	//"math"
)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x IntSet
		//x.Add(rand.Intn(math.MaxInt))
		x.Add(rand.Intn(1000))
	}
}

func BenchmarkUnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x,y IntSet
		x.Add(rand.Intn(1000))
		y.Add(rand.Intn(1000))
		x.UnionWith(&y)
	}
}
