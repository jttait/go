package main

import (
	"fmt"
	"bytes"
)

type IntSet struct {
	words []uint
}

const size = 32 << (^uint(0) >> 63)

func main() {
	var x IntSet
	x.Add(2)
	x.Add(29)
	x.Add(101)
	fmt.Println(x.String())
	fmt.Println(x.Elems())
}

func (s *IntSet) Elems() []int {
	var result []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, size*i+j)
			}
		}
	}
	return result
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/size, uint(x%size)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/size, uint(x%size)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

func (s *IntSet) Remove(x int) {
	word, bit := x/size, uint(x%size)
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	if len(t.words) > len(s.words) {
		t.words = t.words[:len(s.words)]
	}
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i, _ := range s.words {
		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, _ := range s.words {
		if i < len(t.words) {
			s.words[i] &= ^(t.words[i])
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	i := s.Copy()
	s.UnionWith(t)
	i.IntersectWith(t)
	s.DifferenceWith(i)
}

func (s *IntSet) Len() int {
	length := 0
	for _, word := range s.words {
		for bit := 0; bit < size; bit++ {
			if word&(1<<bit) != 0 {
				length++
			}
		}
	}
	return length
}

func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.words = make([]uint, len(s.words))
	copy(t.words, s.words)
	return &t
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", size*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
