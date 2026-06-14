package go_koans

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestKoanRegistrationOrder locks the lesson order and ensures every koan is
// wired to a real function, so future edits to the registry can't silently
// reorder, drop, or forget to register a koan.
func TestKoanRegistrationOrder(t *testing.T) {
	want := []string{
		"basics", "strings", "arrays", "slices", "types",
		"control flow", "enumeration", "anonymous functions",
		"variadic functions", "defer", "files", "interfaces",
		"common interfaces", "maps", "pointers", "structs",
		"allocation", "channels", "concurrency", "panics",
	}

	got := koans()
	if len(got) != len(want) {
		t.Fatalf("koan count = %d, want %d", len(got), len(want))
	}
	for i, name := range want {
		if got[i].name != name {
			t.Errorf("koan[%d] name = %q, want %q", i, got[i].name, name)
		}
		if got[i].run == nil {
			t.Errorf("koan[%d] %q has a nil run func", i, name)
		}
	}
}

// TestRecentLineLocatesCallSite verifies the stack-walking rule assert relies
// on: recentLine(skip) resolves to its caller at the given depth. skip=1 must
// point back at this test's call site.
func TestRecentLineLocatesCallSite(t *testing.T) {
	// recentLine(1) resolves to its own caller. The recentLine(1) statement
	// sits exactly one line above the runtime.Caller(0) call below, so the
	// expected line is one less than what Caller reports here. These two
	// statements must stay adjacent for the offset to hold.
	got := recentLine(1)
	_, file, line, _ := runtime.Caller(0)

	want := fmt.Sprintf("%s:%d", filepath.Base(file), line-1)
	if !strings.HasPrefix(got, want) {
		t.Fatalf("recentLine(1) = %q, want prefix %q", got, want)
	}
	if !strings.Contains(got, "recentLine(1)") {
		t.Errorf("recentLine(1) = %q, want it to include the failing source line", got)
	}
}

// TestFormatFailureLocation covers recent-line extraction end to end against a
// controlled file: basename derivation, 1-based line indexing, whitespace
// trimming, and the "basename:line\ncode" rendering.
func TestFormatFailureLocation(t *testing.T) {
	file := filepath.Join(t.TempDir(), "sample.go")
	contents := "first line\n\tindented assert() // note\nthird line\n"
	if err := ioutil.WriteFile(file, []byte(contents), 0644); err != nil {
		t.Fatal(err)
	}

	got := formatFailureLocation(file, 2)
	want := "sample.go:2\nindented assert() // note"
	if got != want {
		t.Errorf("formatFailureLocation = %q, want %q", got, want)
	}
}

// TestAssertPassesWithoutExit documents that only failed assertions abort the
// run: a satisfied assertion returns normally. If this regressed to exiting,
// the test binary would die here and the test would not pass.
func TestAssertPassesWithoutExit(t *testing.T) {
	assert(true)
}
