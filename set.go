// Copyright 2025 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package ds

import (
	"iter"
)

// A Set represents a mathematical set.
//
// The zero value for Set is usable, and represents an empty set.
// Sets might not work correctly with values which are not comparable in
// the ordinary way, such as floating point NaNs.
type Set[T comparable] struct {
	items map[T]struct{}
}

// init initializes a Set.
func (s *Set[T]) init() {
	if s.items == nil {
		s.items = make(map[T]struct{})
	}
}

// Has returns whether s contains t.
func (s *Set[T]) Has(t T) bool {
	if s.items == nil {
		return false
	}
	_, have := s.items[t]
	return have
}

// Insert adds t to s. If t is already in s, s does not change (we do add a second copy).
func (s *Set[T]) Insert(t T) {
	s.init()
	s.items[t] = struct{}{}
}

// Delete removes t from s.
func (s *Set[T]) Delete(t T) {
	delete(s.items, t)
}

// Len returns the number of items in s.
func (s *Set[T]) Len() int {
	return len(s.items)
}

// Empty returns whether s is an empty set.
func (s *Set[T]) Empty() bool {
	return len(s.items) == 0
}

// All returns an iterator over the elements of s, in arbitrary order.
func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for t := range s.items {
			if !yield(t) {
				return
			}
		}
	}
}
