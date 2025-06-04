// Copyright 2025 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package ds

import (
	"fmt"
	"slices"
	"testing"

	"github.com/pat42smith/gotest"
)

func TestSetBasics(t *testing.T) {
	var set Set[int]
	var bits [10]bool
	for k := 1; k <= 100; k++ {
		d := 1e6 / k % 10
		if k%3 == 0 {
			set.Delete(d)
			bits[d] = false
		} else {
			set.Insert(d)
			bits[d] = true
		}
		ss, sb := "", ""
		for d := 0; d < 10; d++ {
			if set.Has(d) {
				ss = fmt.Sprint(ss, d)
			}
			if bits[d] {
				sb = fmt.Sprint(sb, d)
			}
		}
		gotest.Expect(t, sb, ss)
		gotest.Expect(t, len(sb), set.Len())
	}
}

func TestSetZero(t *testing.T) {
	var s1, s2, s3, s4 Set[string]

	gotest.Expect(t, 0, s1.Len())
	s1.Insert("one")
	gotest.Expect(t, 1, s1.Len())
	s1.Delete("one")
	gotest.Expect(t, 0, s1.Len())

	gotest.Require(t, s2.Empty())
	s2.Insert("two")
	gotest.Require(t, !s2.Empty())
	s2.Delete("two")
	gotest.Require(t, s2.Empty())

	gotest.Require(t, !s3.Has("three"))
	s3.Insert("three")
	gotest.Require(t, s3.Has("three"))
	s3.Delete("three")
	gotest.Require(t, !s3.Has("three"))

	gotest.Require(t, !s4.Has("four"))
	s4.Delete("four")
	gotest.Require(t, !s4.Has("four"))
	gotest.Expect(t, 0, s4.Len())
}

func TestSetAll(t *testing.T) {
	var set Set[string]
	for _ = range set.All() {
		t.Fatal("Set.All on an empty set had an iteration")
	}

	set.Insert("one")
	set.Insert("two")
	set.Insert("three")
	set.Insert("four")
	set.Insert("five")

	var list []string
	for s := range set.All() {
		list = append(list, s)
	}
	slices.Sort(list)

	gotest.Require(t, slices.Equal(list, []string{"five", "four", "one", "three", "two"}))
}
