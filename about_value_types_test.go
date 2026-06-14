package go_koans

import "testing"

// These tests pin the concrete Go behaviour that the value/pointer koans teach
// (about_pointers, about_structs, about_allocation, about_types), independent of
// the __placeholder__ blanks students fill in. They guard against regressions
// when the shared example types (person, coolNumber) are refactored.

func TestValueAssignmentCopies(t *testing.T) {
	a := 3
	b := a // copy, not an alias
	b++

	if a != 3 {
		t.Errorf("assigning a value copies it: original must stay 3, got %d", a)
	}
	if b != 4 {
		t.Errorf("the copy is independent and should change: want 4, got %d", b)
	}
}

func TestPointerMutatesOriginal(t *testing.T) {
	a := 3
	b := &a
	*b = *b + 2
	if a != 5 {
		t.Errorf("dereferencing a pointer mutates the original: want 5, got %d", a)
	}

	byValue := 3
	func(i int) { i++ }(byValue)
	if byValue != 3 {
		t.Errorf("passing by value copies the argument: want 3, got %d", byValue)
	}

	byPointer := 3
	func(i *int) { (*i)++ }(&byPointer)
	if byPointer != 4 {
		t.Errorf("passing a pointer lets the callee mutate the value: want 4, got %d", byPointer)
	}
}

func TestCustomNumberType(t *testing.T) {
	i := coolNumber(4)
	if i != coolNumber(4) {
		t.Errorf("conversion between compatible types preserves the value: want 4, got %d", int(i))
	}
	if got := i.multiplyByTwo(); got != 8 {
		t.Errorf("methods defined on a custom type should run: want 8, got %d", got)
	}
}

func TestStructInitAndComparison(t *testing.T) {
	var john person
	john.name = "bob"
	john.age = 30

	// An anonymous struct with identical fields is comparable to the named
	// person type; this is exactly the equality the structs koan relies on.
	bob := struct {
		name string
		age  int
	}{name: "bob", age: 30}

	if bob != john {
		t.Errorf("structs with equal fields compare equal: bob=%v john=%v", bob, john)
	}

	other := person{name: "alice", age: 31}
	if john == other {
		t.Errorf("structs with differing fields must not be equal: %v vs %v", john, other)
	}
}

func TestNewZeroesValues(t *testing.T) {
	a := new(int)
	if *a != 0 {
		t.Errorf("new(int) points to the zero value: want 0, got %d", *a)
	}
	*a = 3
	if *a != 3 {
		t.Errorf("writing through the pointer sticks: want 3, got %d", *a)
	}

	bob := new(person)
	if bob.name != "" || bob.age != 0 {
		t.Errorf("new(person) allocates a zeroed struct: got %+v", *bob)
	}
}

func TestMakeSlicesAndMaps(t *testing.T) {
	slice := make([]int, 3)
	if len(slice) != 3 {
		t.Errorf("make([]int, 3) sets the length: want 3, got %d", len(slice))
	}

	slice = make([]int, 3, 20)
	if cap(slice) != 20 {
		t.Errorf("make() takes an optional capacity: want 20, got %d", cap(slice))
	}

	m := make(map[int]string)
	if len(m) != 0 {
		t.Errorf("make() builds an empty map: want len 0, got %d", len(m))
	}
}
