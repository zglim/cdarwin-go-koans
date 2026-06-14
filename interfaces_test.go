package go_koans

import (
	"io"
	"testing"
)

// Compile-time proof of the conformance the aboutInterfaces() koan describes:
// no "implements" keyword is needed, the method set alone makes *human and
// *program satisfy the runner interface. If either type stopped implementing
// run(), the package would fail to build right here.
var (
	_ runner = (*human)(nil)
	_ runner = (*program)(nil)
)

func TestRunTwiceAffectsHuman(t *testing.T) {
	h := new(human)
	runTwice(h)
	if h.milesCompleted != 2 {
		t.Fatalf("runTwice should call run() twice: got milesCompleted=%d, want 2", h.milesCompleted)
	}
}

func TestRunTwiceAffectsProgram(t *testing.T) {
	p := new(program)
	runTwice(p)
	if p.executionCount != 2 {
		t.Fatalf("runTwice should call run() twice: got executionCount=%d, want 2", p.executionCount)
	}
}

func TestRunTwiceConcreteTypesAreIndependent(t *testing.T) {
	h := new(human)
	p := new(program)

	runTwice(h) // exercising the *human must not disturb the *program

	if h.milesCompleted != 2 {
		t.Fatalf("running a *human twice: got milesCompleted=%d, want 2", h.milesCompleted)
	}
	if p.executionCount != 0 {
		t.Fatalf("running a *human must not touch a *program: got executionCount=%d, want 0", p.executionCount)
	}
}

func TestRunnerInterfaceInference(t *testing.T) {
	// Passing concrete pointers where a runner is expected is exactly the
	// inference path the koan exercises; the dynamic type stays recoverable.
	var r runner = new(human)
	if _, ok := r.(*human); !ok {
		t.Fatalf("a runner backed by *human should type-assert back to *human")
	}

	r = new(program)
	if _, ok := r.(*program); !ok {
		t.Fatalf("a runner backed by *program should type-assert back to *program")
	}
}

func TestCopyDemoFullCopy(t *testing.T) {
	got := copyDemo("hello world", func(dst io.Writer, src io.Reader) {
		io.Copy(dst, src)
	})
	if got != "hello world" {
		t.Fatalf("io.Copy should move all bytes: got %q, want %q", got, "hello world")
	}
}

func TestCopyDemoTruncatedCopy(t *testing.T) {
	got := copyDemo("hello world", func(dst io.Writer, src io.Reader) {
		io.CopyN(dst, src, 5)
	})
	if got != "hello" {
		t.Fatalf("io.CopyN(.,.,5) should move only the first 5 bytes: got %q, want %q", got, "hello")
	}
}

func TestCopyDemoWithoutCopyLeavesDestinationEmpty(t *testing.T) {
	// Mirrors the unsolved koan state: with no copy operation the io.Writer
	// destination stays empty, which is why the koan asserts currently fail.
	got := copyDemo("hello world", func(dst io.Writer, src io.Reader) {})
	if got != "" {
		t.Fatalf("with no copy the destination must stay empty: got %q, want %q", got, "")
	}
}
