package gotest

import (
	"tang/heap"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	m := heap.MaxHeap{}
	m.Push(1)
	m.Push(3)
	m.Push(2)
	m.Push(5)
	expected := []int{5, 3, 2, 1}
	for _, data := range expected {
		if m.Extract() != data {
			t.Errorf("Error")
		}
	}
}
