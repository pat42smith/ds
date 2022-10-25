// Copyright 2022 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

// Package ds will implement some basic data structures in Go.
//
// For now, only Stack is provided.
package ds

// A Stack is a simple last-in-first-out collection of items.
type Stack[T any] struct {
	items []T
}

// Push inserts a value into a Stack.
func (s *Stack[T]) Push(t T) {
	s.items = append(s.items, t)
}

// Pop removes the value that was most recently inserted into a Stack, and not yet returned by Pop.
//
// It returns the value that was removed.
// Pop panics when called on an empty Stack.
func (s *Stack[T]) Pop() T {
	l := len(s.items)
	t := s.items[l-1]
	s.items = s.items[:l-1]
	return t
}

// Len returns the number of values currently in a Stack.
func (s Stack[T]) Len() int {
	return len(s.items)
}

// Empty returns whether a Stack contains no values.
func (s Stack[T]) Empty() bool {
	return len(s.items) == 0
}

// PushMany inserts several values into a stack.
//
// The values are inserted in the order given, so they will be returned by Pop in the reverse order.
func (s *Stack[T]) PushMany(ts ...T) {
	s.items = append(s.items, ts...)
}

// Peek returns the same value as Pop, but does not remove it from the Stack.
//
// Peek panics when called on an empty Stack.
func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}
