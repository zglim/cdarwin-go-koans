package go_koans

import "testing"

// ---------------------------------------------------------------------------
// Tests for the value-type and pointer koans.
// These ensure the refactoring of about_pointers, about_structs,
// about_allocation, and about_types does not regress teaching behaviour.
// ---------------------------------------------------------------------------

// --- Value vs pointer modification paths ---

func TestValueCopyIsIndependent(t *testing.T) {
	a := 3
	b := a // copy
	b++
	if a != 3 {
		t.Errorf("expected a to remain 3 after copy-and-increment, got %d", a)
	}
}

func TestPointerMutatesOriginal(t *testing.T) {
	a := 3
	b := &a
	*b = *b + 2
	if a != 5 {
		t.Errorf("expected a to become 5 via pointer mutation, got %d", a)
	}
}

func TestPassByValueDoesNotMutate(t *testing.T) {
	increment := func(i int) { i++ }
	a := 3
	increment(a)
	if a != 3 {
		t.Errorf("expected a to remain 3 (pass by value), got %d", a)
	}
}

func TestPassByPointerMutates(t *testing.T) {
	realIncrement := func(i *int) { (*i)++ }
	b := 3
	realIncrement(&b)
	if b != 4 {
		t.Errorf("expected b to become 4 (pass by pointer), got %d", b)
	}
}

// --- Struct initialization and comparison ---

func TestPersonStructFields(t *testing.T) {
	var p person
	p.name = "alice"
	p.age = 25
	if p.name != "alice" {
		t.Errorf("expected name alice, got %s", p.name)
	}
	if p.age != 25 {
		t.Errorf("expected age 25, got %d", p.age)
	}
}

func TestPersonStructEquality(t *testing.T) {
	a := person{name: "bob", age: 30}
	b := person{name: "bob", age: 30}
	c := person{name: "alice", age: 30}

	if a != b {
		t.Error("expected a == b (same field values)")
	}
	if a == c {
		t.Error("expected a != c (different name)")
	}
}

func TestPersonStructLiteralInit(t *testing.T) {
	p := person{name: "john", age: 40}
	if p.name != "john" || p.age != 40 {
		t.Errorf("unexpected person literal: %+v", p)
	}
}

// --- Custom type method ---

func TestCoolNumberMultiplyByTwo(t *testing.T) {
	i := coolNumber(4)
	if i != coolNumber(4) {
		t.Errorf("expected coolNumber(4), got %v", i)
	}
	if i.multiplyByTwo() != 8 {
		t.Errorf("expected multiplyByTwo() == 8, got %d", i.multiplyByTwo())
	}
}

func TestCoolNumberTypeConversion(t *testing.T) {
	var n int = 7
	cn := coolNumber(n)
	if cn.multiplyByTwo() != 14 {
		t.Errorf("expected 14, got %d", cn.multiplyByTwo())
	}
}

// --- new() and make() behaviour ---

func TestNewInt(t *testing.T) {
	a := new(int)
	if *a != 0 {
		t.Errorf("expected new(int) to be zero-valued, got %d", *a)
	}
	*a = 3
	if *a != 3 {
		t.Errorf("expected 3 after assignment, got %d", *a)
	}
}

func TestNewPerson(t *testing.T) {
	bob := new(person)
	// new() returns a pointer; fields are zero-valued.
	if bob.name != "" {
		t.Errorf("expected empty name, got %q", bob.name)
	}
	if bob.age != 0 {
		t.Errorf("expected age 0, got %d", bob.age)
	}
	bob.name = "bob"
	bob.age = 30
	if bob.name != "bob" || bob.age != 30 {
		t.Errorf("unexpected person after mutation: %+v", bob)
	}
}

func TestMakeSliceLength(t *testing.T) {
	slice := make([]int, 3)
	if len(slice) != 3 {
		t.Errorf("expected len 3, got %d", len(slice))
	}
	// make initialises elements to zero value.
	for i, v := range slice {
		if v != 0 {
			t.Errorf("slice[%d] expected 0, got %d", i, v)
		}
	}
}

func TestMakeSliceWithCapacity(t *testing.T) {
	slice := make([]int, 3, 20)
	if len(slice) != 3 {
		t.Errorf("expected len 3, got %d", len(slice))
	}
	if cap(slice) != 20 {
		t.Errorf("expected cap 20, got %d", cap(slice))
	}
}

func TestMakeMap(t *testing.T) {
	m := make(map[int]string)
	if len(m) != 0 {
		t.Errorf("expected empty map, got len %d", len(m))
	}
	m[1] = "one"
	if m[1] != "one" {
		t.Errorf("expected m[1] == 'one', got %q", m[1])
	}
}
