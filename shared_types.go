package go_koans

// Shared types for the value-type and pointer koans
// (about_pointers, about_structs, about_allocation, about_types).
// Centralising them avoids duplicate local definitions while keeping
// the individual koan files focused on the concept they teach.

// person is a simple struct used to demonstrate struct fields,
// value-copy semantics, pointer mutation, and heap allocation.
type person struct {
	name string
	age  int
}

// coolNumber is a user-defined numeric type used to demonstrate
// custom type definitions and method receivers.
type coolNumber int

// multiplyByTwo returns the coolNumber value multiplied by 2.
func (cn coolNumber) multiplyByTwo() int {
	return int(cn) * 2
}
