package go_koans

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

const (
	__string__       string  = "impossibly lame value"
	__int__          int     = -1
	__positive_int__ int     = 42
	__byte__         byte    = 255
	__bool__         bool    = false // ugh
	__boolean__      bool    = true  // oh well
	__float32__      float32 = -1.0
	__delete_me__    bool    = false
)

var __runner__ runner = nil

// koan pairs a human-readable lesson name with the function that runs it.
// The name is metadata only: it lets the registry document itself and lets
// tests lock the ordering. It is never printed during a normal run, so the
// student-facing output is unchanged.
type koan struct {
	name string
	run  func()
}

// koans is the single registry of lessons, listed in execution order.
// Adding a new koan file means appending its about* function here once — no
// other execution code needs to change.
func koans() []koan {
	return []koan{
		{"basics", aboutBasics},
		{"strings", aboutStrings},
		{"arrays", aboutArrays},
		{"slices", aboutSlices},
		{"types", aboutTypes},
		{"control flow", aboutControlFlow},
		{"enumeration", aboutEnumeration},
		{"anonymous functions", aboutAnonymousFunctions},
		{"variadic functions", aboutVariadicFunctions},
		{"defer", aboutDefer},
		{"files", aboutFiles},
		{"interfaces", aboutInterfaces},
		{"common interfaces", aboutCommonInterfaces},
		{"maps", aboutMaps},
		{"pointers", aboutPointers},
		{"structs", aboutStructs},
		{"allocation", aboutAllocation},
		{"channels", aboutChannels},
		{"concurrency", aboutConcurrency},
		{"panics", aboutPanics},
	}
}

func TestKoans(t *testing.T) {
	runKoans(koans())
	celebrate()
}

// runKoans executes each registered koan in order. A failing assert inside a
// koan ends the process, so reaching the end means every lesson passed.
func runKoans(ks []koan) {
	for _, k := range ks {
		k.run()
	}
}

// celebrate prints the success banner shown once all koans pass.
func celebrate() {
	fmt.Printf("\n%c[32;1mYou won life. Good job.%c[0m\n\n", 27, 27)
}

// assert stops the whole run at the first unmet expectation, pointing the
// student at the exact source line that failed.
func assert(ok bool) {
	if !ok {
		reportFailure(recentLine(2))
	}
}

// reportFailure renders the failing location in magenta and aborts the run.
// It owns the output-and-exit step so the location lookup stays pure.
func reportFailure(location string) {
	fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, location, 27)
	os.Exit(1)
}

// recentLine resolves the source location `skip` frames up the call stack,
// following runtime.Caller semantics, and formats it for display. assert
// passes skip=2 so the reported line is the koan's assert call rather than
// assert itself.
func recentLine(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return formatFailureLocation(file, line)
}

// formatFailureLocation reads `file`, pulls out the 1-based `line`, and renders
// "basename:line" followed by the trimmed source of that line.
func formatFailureLocation(file string, line int) string {
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
