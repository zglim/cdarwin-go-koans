package go_koans

import (
	"fmt"
	"testing"
)

// These tests lock in the answers the collection koans expect, proving that
// routing their example data through the shared helpers in collections_examples.go
// does not change what arrays, slices, enumeration, or variadic functions
// observe. They deliberately re-create each koan's operations instead of calling
// the aboutX() functions, because those still contain unsolved placeholder
// blanks and call os.Exit on the first failed assertion.
//
// Run them in isolation with:  GO111MODULE=off go test -run TestCollections

func TestCollectionsFruitNamesShared(t *testing.T) {
	want := []string{"apple", "orange", "mango"}
	got := fruitNames()
	if len(got) != len(want) {
		t.Fatalf("fruitNames() length = %d, want %d", len(got), len(want))
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("fruitNames()[%d] = %q, want %q", i, got[i], w)
		}
	}

	// Each call must return a fresh slice so a mutation in one koan cannot leak
	// into another.
	a := fruitNames()
	a[0] = "lemon"
	if fruitNames()[0] != "apple" {
		t.Error("fruitNames() shares backing data between calls; mutation leaked")
	}
}

func TestCollectionsFruitArray(t *testing.T) {
	fruits := fruitArray()

	if got := fmt.Sprintf("%T", fruits); got != "[4]string" {
		t.Errorf("fruitArray() type = %s, want [4]string", got)
	}
	if fruits[0] != "apple" || fruits[1] != "orange" || fruits[2] != "mango" {
		t.Errorf("fruitArray() values = %v, want [apple orange mango]", fruits)
	}
	if fruits[3] != "" {
		t.Errorf("fruitArray()[3] = %q, want empty trailing zero value", fruits[3])
	}
	if len(fruits) != 4 {
		t.Errorf("len(fruitArray()) = %d, want 4", len(fruits))
	}
	if cap(fruits) != 4 {
		t.Errorf("cap(fruitArray()) = %d, want 4", cap(fruits))
	}

	// Slicing an array yields a []string that shares the backing data.
	tasty := fruits[1:3]
	if got := fmt.Sprintf("%T", tasty); got != "[]string" {
		t.Errorf("slice of array type = %s, want []string", got)
	}
	if tasty[0] != "orange" || tasty[1] != "mango" {
		t.Errorf("fruits[1:3] = %v, want [orange mango]", tasty)
	}
	if len(tasty) != 2 {
		t.Errorf("len(fruits[1:3]) = %d, want 2", len(tasty))
	}
	if cap(tasty) != 3 {
		t.Errorf("cap(fruits[1:3]) = %d, want 3", cap(tasty))
	}

	// Mutating the slice writes through to the array, just as the koan teaches.
	tasty[0] = "lemon"
	if fruits[1] != "lemon" {
		t.Errorf("array not updated through slice: fruits[1] = %q, want lemon", fruits[1])
	}
	if fruits[0] != "apple" || fruits[2] != "mango" || fruits[3] != "" {
		t.Errorf("unexpected array mutation: %v", fruits)
	}
}

func TestCollectionsFruitSlice(t *testing.T) {
	fruits := fruitNames()

	if got := fmt.Sprintf("%T", fruits); got != "[]string" {
		t.Errorf("fruitNames() type = %s, want []string", got)
	}
	if fruits[0] != "apple" {
		t.Errorf("fruits[0] = %q, want apple", fruits[0])
	}
	if len(fruits) != 3 {
		t.Errorf("len(fruits) = %d, want 3", len(fruits))
	}

	tasty := fruits[1:3]
	if tasty[0] != "orange" {
		t.Errorf("fruits[1:3][0] = %q, want orange", tasty[0])
	}
}

func TestCollectionsSliceAppendGrowth(t *testing.T) {
	// Mirror the slice koan's append sequence to lock in its len progression.
	slots := []string{"baby", "baby", "lemon"}
	if cap(slots) != 3 {
		t.Errorf("initial cap = %d, want 3 (capacity starts as the literal length)", cap(slots))
	}

	slots = append(slots, "baby!")
	if len(slots) != 4 {
		t.Errorf("len after single append = %d, want 4", len(slots))
	}

	slots = append(slots, "another baby!?", "yet another, oh dear!", "they must be Catholic")
	if len(slots) != 7 {
		t.Errorf("len after batch append = %d, want 7", len(slots))
	}

	// Capacity growth is allocator-defined, so assert it matches an independent
	// reference grown the same way rather than a hard-coded number. This keeps
	// the koan's "guessable algorithm" expectation valid across Go versions.
	ref := []string{"baby", "baby", "lemon"}
	ref = append(ref, "baby!")
	ref = append(ref, "another baby!?", "yet another, oh dear!", "they must be Catholic")
	if cap(slots) != cap(ref) {
		t.Errorf("cap drifted from reference: got %d, want %d", cap(slots), cap(ref))
	}
}

func TestCollectionsEnumeration(t *testing.T) {
	want := []string{"hello", " world", "!"}
	fragments := greetingFragments()
	for i, w := range want {
		if fragments[i] != w {
			t.Errorf("greetingFragments()[%d] = %q, want %q", i, fragments[i], w)
		}
	}

	// Block 1: range exposes both index and value.
	var concatenated string
	var total int
	for i, v := range fragments {
		total += i
		concatenated += v
	}
	if concatenated != "hello world!" {
		t.Errorf("concatenated = %q, want %q", concatenated, "hello world!")
	}
	if total != 3 {
		t.Errorf("total = %d, want 3", total)
	}

	// Block 2: range with the index omitted, summing value lengths.
	var totalLength int
	for _, v := range greetingFragments() {
		totalLength += len(v)
	}
	if totalLength != 12 {
		t.Errorf("totalLength = %d, want 12", totalLength)
	}
}

func TestCollectionsVariadic(t *testing.T) {
	want := []string{"bob", "billy", "fred"}
	friends := friendNames()
	for i, w := range want {
		if friends[i] != w {
			t.Errorf("friendNames()[%d] = %q, want %q", i, friends[i], w)
		}
	}

	// Discrete arguments, exactly as the first koan block passes them.
	if got := concatNames(" ", friends[0], friends[1], friends[2]); got != "bob billy fred" {
		t.Errorf("concatNames with discrete args = %q, want %q", got, "bob billy fred")
	}

	// A slice spread with ..., exactly as the second koan block passes it.
	if got := concatNames("-", friends...); got != "bob-billy-fred" {
		t.Errorf("concatNames with spread slice = %q, want %q", got, "bob-billy-fred")
	}
}
