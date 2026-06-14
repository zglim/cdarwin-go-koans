package go_koans

// ---------------------------------------------------------------------------
// Shared example data for collection-related koans.
//
// These variables and small helpers are the single source of truth for the
// string lists that appear in about_arrays, about_slices, about_enumeration
// and about_variadic_functions.  Each koan file builds on top of them so that
// renaming a value or adjusting an example only needs to happen here.
// ---------------------------------------------------------------------------

// sampleFruits is the canonical string list used wherever the koans need a
// short, fixed set of fruit names (arrays, slices, and their derivatives).
var sampleFruits = []string{"apple", "orange", "mango"}

// sampleNames is the canonical list of person names used by the variadic
// function koans.
var sampleNames = []string{"bob", "billy", "fred"}

// sampleGreetings is the canonical list used by the enumeration koans to
// demonstrate range-based iteration and string concatenation.
var sampleGreetings = []string{"hello", " world", "!"}

// sampleSlots is the canonical list used to demonstrate append and capacity
// growth in the slices koans.
var sampleSlots = []string{"baby", "baby", "lemon"}

// extraSlotItems are the additional items appended in bulk to demonstrate
// variadic append behaviour on slices.
var extraSlotItems = []string{"another baby!?", "yet another, oh dear!", "they must be Catholic"}

// sampleVeggies is the canonical short list used to show that array literals
// can infer their length from the element count.
var sampleVeggies = [...]string{"carrot", "pea"}

// --- small helpers ---------------------------------------------------------

// concatenateAll joins every element of ss into a single string (no
// separator).  Used by the enumeration koans to verify range accumulation.
func concatenateAll(ss []string) string {
	var out string
	for _, v := range ss {
		out += v
	}
	return out
}

// sumIndices returns the sum of all indices produced by ranging over ss.
func sumIndices(ss []string) int {
	total := 0
	for i := range ss {
		total += i
	}
	return total
}

// totalLength returns the sum of len(s) for every string in ss.
func totalLength(ss []string) int {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	return n
}
