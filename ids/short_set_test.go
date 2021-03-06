// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ids

import (
	"strings"
	"testing"
)

func TestShortSetContains(t *testing.T) {
	set := ShortSet{}

	id0 := NewShortID([20]byte{0})
	id1 := NewShortID([20]byte{1})

	switch {
	case set.Contains(id0):
		t.Fatalf("Sets shouldn't contain %s", id0)
	case set.Contains(id1):
		t.Fatalf("Sets shouldn't contain %s", id1)
	}

	set.Add(id0)

	switch {
	case !set.Contains(id0):
		t.Fatalf("Set should contain %s", id0)
	case set.Contains(id1):
		t.Fatalf("Set shouldn't contain %s", id1)
	}

	set.Add(id1)

	switch {
	case !set.Contains(id0):
		t.Fatalf("Set should contain %s", id0)
	case !set.Contains(id1):
		t.Fatalf("Set should contain %s", id1)
	}

	set.Remove(id0)

	switch {
	case set.Contains(id0):
		t.Fatalf("Sets shouldn't contain %s", id0)
	case !set.Contains(id1):
		t.Fatalf("Set should contain %s", id1)
	}

	set.Add(id0)

	switch {
	case !set.Contains(id0):
		t.Fatalf("Set should contain %s", id0)
	case !set.Contains(id1):
		t.Fatalf("Set should contain %s", id1)
	}
}

func TestShortSetUnion(t *testing.T) {
	set := ShortSet{}
	unionSet := ShortSet{}

	id0 := NewShortID([20]byte{0})
	id1 := NewShortID([20]byte{1})

	unionSet.Add(id0)
	set.Union(unionSet)

	switch {
	case !set.Contains(id0):
		t.Fatalf("Set should contain %s", id0)
	case set.Contains(id1):
		t.Fatalf("Set shouldn't contain %s", id1)
	}

	unionSet.Add(id1)
	set.Union(unionSet)

	switch {
	case !set.Contains(id0):
		t.Fatalf("Set should contain %s", id0)
	case !set.Contains(id1):
		t.Fatalf("Set should contain %s", id1)
	}

	set.Remove(id0)

	switch {
	case set.Contains(id0):
		t.Fatalf("Sets shouldn't contain %s", id0)
	case !set.Contains(id1):
		t.Fatalf("Set should contain %s", id1)
	}

	set.Clear()
	set.Union(unionSet)

	switch {
	case !set.Contains(id0):
		t.Fatalf("Set should contain %s", id0)
	case !set.Contains(id1):
		t.Fatalf("Set should contain %s", id1)
	}
}

func TestShortSetEquals(t *testing.T) {
	set := ShortSet{}
	otherSet := ShortSet{}
	if !set.Equals(otherSet) {
		t.Fatal("Empty sets should be equal")
	}
	if !otherSet.Equals(set) {
		t.Fatal("Empty sets should be equal")
	}

	set.Add(NewShortID([20]byte{1, 2, 3, 4, 5}))
	if set.Equals(otherSet) {
		t.Fatal("Sets should be unequal")
	}
	if otherSet.Equals(set) {
		t.Fatal("Sets should be unequal")
	}

	otherSet.Add(NewShortID([20]byte{1, 2, 3, 4, 5}))
	if !set.Equals(otherSet) {
		t.Fatal("sets should be equal")
	}
	if !otherSet.Equals(set) {
		t.Fatal("sets should be equal")
	}

	otherSet.Add(NewShortID([20]byte{6, 7, 8, 9, 10}))
	if set.Equals(otherSet) {
		t.Fatal("Sets should be unequal")
	}
	if otherSet.Equals(set) {
		t.Fatal("Sets should be unequal")
	}

	set.Add(NewShortID([20]byte{6, 7, 8, 9, 10}))
	if !set.Equals(otherSet) {
		t.Fatal("sets should be equal")
	}
	if !otherSet.Equals(set) {
		t.Fatal("sets should be equal")
	}

	otherSet.Add(NewShortID([20]byte{11, 12, 13, 14, 15}))
	if set.Equals(otherSet) {
		t.Fatal("Sets should be unequal")
	}
	if otherSet.Equals(set) {
		t.Fatal("Sets should be unequal")
	}

	set.Add(NewShortID([20]byte{11, 12, 13, 14, 16}))
	if set.Equals(otherSet) {
		t.Fatal("Sets should be unequal")
	}
	if otherSet.Equals(set) {
		t.Fatal("Sets should be unequal")
	}
}

func TestShortSetList(t *testing.T) {
	set := ShortSet{}
	otherSet := ShortSet{}

	id0 := NewShortID([20]byte{0})
	id1 := NewShortID([20]byte{1})

	set.Add(id0)
	otherSet.Add(set.List()...)

	if !set.Equals(otherSet) {
		t.Fatalf("Sets should be equal but are:\n%s\n%s", set, otherSet)
	}

	set.Add(id1)
	otherSet.Clear()
	otherSet.Add(set.List()...)

	if !set.Equals(otherSet) {
		t.Fatalf("Sets should be equal but are:\n%s\n%s", set, otherSet)
	}
}

func TestShortSetCappedList(t *testing.T) {
	set := ShortSet{}

	id := ShortEmpty

	if list := set.CappedList(0); len(list) != 0 {
		t.Fatalf("List should have been empty but was %v", list)
	}

	set.Add(id)

	if list := set.CappedList(0); len(list) != 0 {
		t.Fatalf("List should have been empty but was %v", list)
	} else if list := set.CappedList(1); len(list) != 1 {
		t.Fatalf("List should have had length %d but had %d", 1, len(list))
	} else if returnedID := list[0]; !id.Equals(returnedID) {
		t.Fatalf("List should have been %s but was %s", id, returnedID)
	} else if list := set.CappedList(2); len(list) != 1 {
		t.Fatalf("List should have had length %d but had %d", 1, len(list))
	} else if returnedID := list[0]; !id.Equals(returnedID) {
		t.Fatalf("List should have been %s but was %s", id, returnedID)
	}
}

func TestShortSetString(t *testing.T) {
	set := ShortSet{}

	id0 := NewShortID([20]byte{0})
	id1 := NewShortID([20]byte{1})

	if str := set.String(); str != "{}" {
		t.Fatalf("Set should have been %s but was %s", "{}", str)
	}

	set.Add(id0)

	if str := set.String(); str != "{111111111111111111116DBWJs}" {
		t.Fatalf("Set should have been %s but was %s", "{111111111111111111116DBWJs}", str)
	}

	set.Add(id1)

	if str := set.String(); !strings.Contains(str, "111111111111111111116DBWJs") {
		t.Fatalf("Set should have contained %s", "111111111111111111116DBWJs")
	} else if !strings.Contains(str, "6HgC8KRBEhXYbF4riJyJFLSHt37UNuRt") {
		t.Fatalf("Set should have contained %s", "6HgC8KRBEhXYbF4riJyJFLSHt37UNuRt")
	} else if count := strings.Count(str, ","); count != 1 {
		t.Fatalf("Should only have one %s in %s", ",", str)
	}
}
