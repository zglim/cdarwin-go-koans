package go_koans

func aboutSlices() {
	// Re-use the shared fruit list so examples stay consistent with the
	// arrays and enumeration koans.
	fruits := make([]string, len(sampleFruits))
	copy(fruits, sampleFruits)

	assert(fruits[0] == __string__) // slices seem like arrays
	assert(len(fruits) == __int__)  // in nearly all respects

	tastyFruits := fruits[1:3]          // we can even slice slices
	assert(tastyFruits[0] == __string__) // slices of slices also share the underlying data

	slots := make([]string, len(sampleSlots))
	copy(slots, sampleSlots)
	assert(cap(slots) == __int__) // the capacity is initially the length

	slots = append(slots, "baby!")
	assert(len(slots) == __int__) // slices can be extended with append(), much like realloc in C
	assert(cap(slots) == __int__) // but with better optimizations

	slots = append(slots, extraSlotItems...)

	assert(len(slots) == __int__) // append() can take N arguments to append to the slice
	assert(cap(slots) == __int__) // the capacity optimizations have a guessable algorithm
}
