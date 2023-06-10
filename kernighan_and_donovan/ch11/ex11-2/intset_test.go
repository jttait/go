package intset

import "testing"

func TestAdd(t *testing.T) {
	var tests = []struct{
		input []int
		want map[int]struct{}
	}{
		{[]int{1}, map[int]struct{}{1: struct{}{}}},
		{[]int{1,1}, map[int]struct{}{1: struct{}{}}},
		{[]int{1,2}, map[int]struct{}{1: struct{}{}, 2: struct{}{}}},
	}
	for _, test := range tests {
		var got IntSet
		for _, i := range test.input {
			got.Add(i)
		}
		for i := range test.want {
			if _, ok := test.want[i]; ok != got.Has(i) {
				t.Error(`Add()`)
			}
		}

	}
}

func TestUnionWith(t *testing.T) {
	var tests = []struct{
		input1 []int
		input2 []int
		want map[int]struct{}
	}{
		{[]int{1}, []int{2}, map[int]struct{}{1: struct{}{}, 2: struct{}{}}},
		{[]int{1}, []int{1}, map[int]struct{}{1: struct{}{}}},
		{[]int{3}, []int{3}, map[int]struct{}{3: struct{}{}}},
	}
	for _, test := range tests {
		var intSet1 IntSet
		var intSet2 IntSet
		for _, i := range test.input1 {
			intSet1.Add(i)
		}
		for _, i := range test.input2 {
			intSet2.Add(i)
		}
		intSet1.UnionWith(&intSet2)
		for i := range test.want {
			if _, ok := test.want[i]; ok != intSet1.Has(i) {
				t.Error(`Add()`)
			}
		}
	}
}
