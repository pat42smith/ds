// Copyright 2022 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

// Package ds will implement some basic data structures in Go.
//
// For now, only Stack is provided.
package ds

import (
	"testing"

	"github.com/pat42smith/gotest"
)

func TestThrees(t *testing.T) {
	var s Stack[int]
	gotest.Expect(t, 0, s.Len())
	gotest.Require(t, s.Empty())

	for n := 0; n < 30; n += 3 {
		s.Push(n)
		gotest.Expect(t, n/3+1, s.Len())
		gotest.Require(t, !s.Empty())

		s.Push(n + 1)
		gotest.Expect(t, n/3+2, s.Len())
		gotest.Require(t, !s.Empty())

		gotest.Expect(t, n+1, s.Pop())
		gotest.Expect(t, n/3+1, s.Len())
		gotest.Require(t, !s.Empty())
	}

	for n := 27; n >= 0; n -= 3 {
		gotest.Expect(t, n, s.Pop())
	}

	gotest.Expect(t, 0, s.Len())
	gotest.Require(t, s.Empty())
	gotest.MustPanic(t, func() { s.Pop() })
}

func TestPushMany(t *testing.T) {
	var s Stack[string]
	s.PushMany("alpha", "beta", "gamma", "delta")
	s.PushMany()
	s.PushMany("un", "deux", "trois")
	gotest.Expect(t, "trois", s.Pop())
	gotest.Expect(t, "deux", s.Pop())
	gotest.Expect(t, "un", s.Pop())
	gotest.Expect(t, "delta", s.Pop())
	gotest.Expect(t, "gamma", s.Pop())
	gotest.Expect(t, "beta", s.Pop())
	gotest.Expect(t, "alpha", s.Pop())
}
