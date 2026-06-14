package go_koans

func aboutTypes() {
	// coolNumber is defined in shared_types.go.
	i := coolNumber(4)
	assert(i == coolNumber(__int__))     // values can be converted between compatible types
	assert(i.multiplyByTwo() == __int__) // you can add methods on any type you define
}
