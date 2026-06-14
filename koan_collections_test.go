package go_koans

import "testing"

// TestSharedDataValues verifies that the shared example data in
// koan_collections.go contains exactly the values that the original,
// pre-refactor koans hard-coded.  If a future edit accidentally changes
// a value here, these assertions catch it before the koans are affected.
func TestSharedDataValues(t *testing.T) {
	// --- sampleFruits (used by about_arrays and about_slices) ----------
	expectedFruits := []string{"apple", "orange", "mango"}
	if len(sampleFruits) != len(expectedFruits) {
		t.Fatalf("sampleFruits length: got %d, want %d", len(sampleFruits), len(expectedFruits))
	}
	for i, v := range expectedFruits {
		if sampleFruits[i] != v {
			t.Errorf("sampleFruits[%d] = %q, want %q", i, sampleFruits[i], v)
		}
	}

	// Verify that copy into [4]string reproduces the original array
	// {"apple", "orange", "mango", ""}.
	var arr [4]string
	copy(arr[:], sampleFruits)
	if arr != [4]string{"apple", "orange", "mango", ""} {
		t.Errorf("array from sampleFruits = %v, want {apple orange mango \"\"}", arr)
	}

	// --- sampleNames (used by about_variadic_functions) ----------------
	expectedNames := []string{"bob", "billy", "fred"}
	if len(sampleNames) != len(expectedNames) {
		t.Fatalf("sampleNames length: got %d, want %d", len(sampleNames), len(expectedNames))
	}
	for i, v := range expectedNames {
		if sampleNames[i] != v {
			t.Errorf("sampleNames[%d] = %q, want %q", i, sampleNames[i], v)
		}
	}

	// --- sampleGreetings (used by about_enumeration) -------------------
	expectedGreetings := []string{"hello", " world", "!"}
	if len(sampleGreetings) != len(expectedGreetings) {
		t.Fatalf("sampleGreetings length: got %d, want %d", len(sampleGreetings), len(expectedGreetings))
	}
	for i, v := range expectedGreetings {
		if sampleGreetings[i] != v {
			t.Errorf("sampleGreetings[%d] = %q, want %q", i, sampleGreetings[i], v)
		}
	}

	// --- sampleSlots (used by about_slices) ----------------------------
	expectedSlots := []string{"baby", "baby", "lemon"}
	if len(sampleSlots) != len(expectedSlots) {
		t.Fatalf("sampleSlots length: got %d, want %d", len(sampleSlots), len(expectedSlots))
	}
	for i, v := range expectedSlots {
		if sampleSlots[i] != v {
			t.Errorf("sampleSlots[%d] = %q, want %q", i, sampleSlots[i], v)
		}
	}

	// --- extraSlotItems ------------------------------------------------
	expectedExtra := []string{"another baby!?", "yet another, oh dear!", "they must be Catholic"}
	if len(extraSlotItems) != len(expectedExtra) {
		t.Fatalf("extraSlotItems length: got %d, want %d", len(extraSlotItems), len(expectedExtra))
	}
	for i, v := range expectedExtra {
		if extraSlotItems[i] != v {
			t.Errorf("extraSlotItems[%d] = %q, want %q", i, extraSlotItems[i], v)
		}
	}

	// --- sampleVeggies -------------------------------------------------
	if len(sampleVeggies) != 2 {
		t.Errorf("len(sampleVeggies) = %d, want 2", len(sampleVeggies))
	}
	if sampleVeggies[0] != "carrot" || sampleVeggies[1] != "pea" {
		t.Errorf("sampleVeggies = %v, want [carrot pea]", sampleVeggies)
	}
}

// TestSharedHelpers verifies the helper functions in koan_collections.go.
func TestSharedHelpers(t *testing.T) {
	// concatenateAll
	got := concatenateAll(sampleGreetings)
	if got != "hello world!" {
		t.Errorf("concatenateAll(sampleGreetings) = %q, want %q", got, "hello world!")
	}

	// sumIndices: for a 3-element slice, indices are 0+1+2 = 3
	gotSum := sumIndices(sampleGreetings)
	if gotSum != 3 {
		t.Errorf("sumIndices(sampleGreetings) = %d, want 3", gotSum)
	}

	// totalLength: len("hello") + len(" world") + len("!") = 5+6+1 = 12
	gotLen := totalLength(sampleGreetings)
	if gotLen != 12 {
		t.Errorf("totalLength(sampleGreetings) = %d, want 12", gotLen)
	}
}

// TestSharedDataSliceAppend verifies that appending to a copy of
// sampleSlots reproduces the length/capacity behaviour the slices koan
// asserts.  This catches any drift in the underlying data that would
// silently change the capacity-growth results.
func TestSharedDataSliceAppend(t *testing.T) {
	slots := make([]string, len(sampleSlots))
	copy(slots, sampleSlots)

	// Initial capacity should equal length (3).
	if cap(slots) != 3 {
		t.Errorf("initial cap = %d, want 3", cap(slots))
	}

	// After appending one element the Go runtime doubles capacity.
	slots = append(slots, "baby!")
	if len(slots) != 4 {
		t.Errorf("after 1 append: len = %d, want 4", len(slots))
	}
	if cap(slots) != 6 {
		t.Errorf("after 1 append: cap = %d, want 6", cap(slots))
	}

	// Append three more items (bulk append via spread).
	slots = append(slots, extraSlotItems...)
	if len(slots) != 7 {
		t.Errorf("after bulk append: len = %d, want 7", len(slots))
	}
	// Capacity doubles again from 6 to 12 when 7 > 6.
	if cap(slots) != 12 {
		t.Errorf("after bulk append: cap = %d, want 12", cap(slots))
	}
}

// TestConcatNamesHelper verifies the concatNames function used by the
// variadic koans, using the shared sampleNames data.
func TestConcatNamesHelper(t *testing.T) {
	// Individual arguments (mirrors the first koan block).
	got := concatNames(" ", sampleNames[0], sampleNames[1], sampleNames[2])
	if got != "bob billy fred" {
		t.Errorf("concatNames individual = %q, want %q", got, "bob billy fred")
	}

	// Slice spread (mirrors the second koan block).
	got = concatNames("-", sampleNames...)
	if got != "bob-billy-fred" {
		t.Errorf("concatNames spread = %q, want %q", got, "bob-billy-fred")
	}
}

// TestSharedDataCrossConsistency ensures that the arrays koan's
// tastyFruits slice and the slices koan's tastyFruits slice start from
// the same underlying data.
func TestSharedDataCrossConsistency(t *testing.T) {
	// Arrays koan: build a [4]string and slice it at [1:3].
	var arr [4]string
	copy(arr[:], sampleFruits)
	arraySlice := arr[1:3]

	// Slices koan: copy sampleFruits into a new slice and slice at [1:3].
	sliceCopy := make([]string, len(sampleFruits))
	copy(sliceCopy, sampleFruits)
	sliceSlice := sliceCopy[1:3]

	// Both should yield the same first element (orange).
	if arraySlice[0] != sliceSlice[0] {
		t.Errorf("arrays tastyFruits[0] = %q != slices tastyFruits[0] = %q", arraySlice[0], sliceSlice[0])
	}
	if arraySlice[1] != sliceSlice[1] {
		t.Errorf("arrays tastyFruits[1] = %q != slices tastyFruits[1] = %q", arraySlice[1], sliceSlice[1])
	}
}
