package go_koans

func aboutAllocation() {
	a := new(int)
	*a = 3
	assert(*a == __int__) // new() creates a pointer to the given type, like malloc() in C

	bob := new(person)         // person is the shared struct from the structs koan
	assert(bob.age == __int__) // it can allocate memory for custom types as well

	slice := make([]int, 3)
	assert(len(slice) == __int__) // make() creates slices of a given length

	slice = make([]int, 3, __positive_int__) // but can also take an optional capacity
	assert(cap(slice) == 20)

	m := make(map[int]string)
	assert(len(m) == __int__) // make() also creates maps
}
