package go_koans

// The collection koans (arrays, slices, enumeration, and variadic functions)
// all lean on a few small, fixed example datasets. Centralising them here keeps
// each koan focused on the behaviour it teaches instead of re-typing the same
// string lists. Every helper returns a fresh value, so a koan that mutates its
// data (via append or index assignment) never disturbs another koan or a re-run.

// fruitNames are the canonical fruits shared by the array and slice koans. Only
// three names are listed on purpose: the array koan packs them into a larger
// [4]string so the trailing zero value stays observable.
func fruitNames() []string {
	return []string{"apple", "orange", "mango"}
}

// fruitArray packs fruitNames into a fixed-size array. The [4]string return type
// keeps the "an array literal may set fewer elements than its length" lesson
// intact while sourcing the values from a single place.
func fruitArray() [4]string {
	var fruits [4]string
	copy(fruits[:], fruitNames())
	return fruits
}

// greetingFragments are the pieces the enumeration koan folds together while
// demonstrating range over a slice.
func greetingFragments() []string {
	return []string{"hello", " world", "!"}
}

// friendNames is the single source for the names the variadic koan passes both
// as discrete arguments and as a spread slice.
func friendNames() []string {
	return []string{"bob", "billy", "fred"}
}
