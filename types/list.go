package types

import (
	"fmt"
	"reflect"
)

type List[T any] struct {
	Data []T
}

func NewList[T any]() *List[T] {
	return &List[T]{
		Data: []T{},
	}
}

// Get the value at the given index.
// If the index is higher than the length, a panic will be thrown.
//
// Args:
// - index: the index of the value to get.
//
// Returns:
// - T: the value at the given index.
func (l *List[T]) Get(index int) T {
	if index > len(l.Data)-1 {
		err := fmt.Sprintf("the given index (%d) is higher than the length (%d)", index, len(l.Data))
		panic(err)
	}
	return l.Data[index]
}

// Insert will add a new value to the list.
//
// Args:
// - v: the value to add to the list.
func (l *List[T]) Insert(v T) {
	l.Data = append(l.Data, v)
}

// Clear will remove all values from the list.
func (l *List[T]) Clear() {
	l.Data = []T{}
}

// GetIndex will return the index of v.
// If v does not exist in the list -1 will be returned.
//
// Args:
// - v: the value to get the index of.
//
// Returns:
// - int: the index of v.
func (l *List[T]) GetIndex(v T) int {
	for i := 0; i < l.Len(); i++ {
		if reflect.DeepEqual(v, l.Data[i]) {
			return i
		}
	}
	return -1
}

// Remove will remove the first occurrence of v from the list.
//
// Args:
// - v: the value to remove from the list.
func (l *List[T]) Remove(v T) {
	index := l.GetIndex(v)
	if index == -1 {
		return
	}
	l.Pop(index)
}

// Pop will remove the value at the given index.
//
// Args:
// - index: the index of the value to remove.
func (l *List[T]) Pop(index int) {
	l.Data = append(l.Data[:index], l.Data[index+1:]...)
}

// Contains will return true if the list contains v.
//
// Args:
// - v: the value to check if it exists in the list.
//
// Returns:
// - bool: true if the list contains v, false otherwise.
func (l *List[T]) Contains(v T) bool {
	for i := 0; i < len(l.Data); i++ {
		if reflect.DeepEqual(l.Data[i], v) {
			return true
		}
	}
	return false
}

func (l List[T]) Last() T {
	return l.Data[l.Len()-1]
}

func (l *List[T]) Len() int {
	return len(l.Data)
}
