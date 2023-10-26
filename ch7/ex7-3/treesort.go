package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	value int
	left, right *tree
}

func (t *tree) String() string {
	result := preOrderTraversal(t, "", 0)
	return result
}

func preOrderTraversal(t *tree, result string, depth int) string {
	for i := 0; i < depth; i++ {
		result += "   "
	}
	if t == nil {
		result += "nil\n"
		return result
	}
	result += strconv.Itoa(t.value) + "\n"
	result = preOrderTraversal(t.left, result, depth+1)
	result = preOrderTraversal(t.right, result, depth+1)
	return result
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	fmt.Println(root.String())
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	integers := []int{1, 200, 62, 10, 110}
	Sort(integers)
	fmt.Println(integers)
}
