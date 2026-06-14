package go_koans

import (
	"io"
	"testing"
)

// ---- runTwice behavior ----

func TestRunTwiceHuman(t *testing.T) {
	h := new(human)
	if h.milesCompleted != 0 {
		t.Fatalf("expected initial milesCompleted to be 0, got %d", h.milesCompleted)
	}
	runTwice(h)
	if h.milesCompleted != 2 {
		t.Errorf("expected milesCompleted to be 2 after runTwice, got %d", h.milesCompleted)
	}
}

func TestRunTwiceProgram(t *testing.T) {
	p := new(program)
	if p.executionCount != 0 {
		t.Fatalf("expected initial executionCount to be 0, got %d", p.executionCount)
	}
	runTwice(p)
	if p.executionCount != 2 {
		t.Errorf("expected executionCount to be 2 after runTwice, got %d", p.executionCount)
	}
}

// ---- interface implementation inference ----

func TestInterfaceInference(t *testing.T) {
	bob := new(human)
	rspec := new(program)

	// Both *human and *program implicitly satisfy runner — no explicit declaration needed.
	var r runner

	r = bob
	r.run()
	if bob.milesCompleted != 1 {
		t.Errorf("expected human milesCompleted to be 1 after one run, got %d", bob.milesCompleted)
	}

	r = rspec
	r.run()
	if rspec.executionCount != 1 {
		t.Errorf("expected program executionCount to be 1 after one run, got %d", rspec.executionCount)
	}

	// A runner holding a concrete value is not nil.
	if r == nil {
		t.Error("runner should not be nil after assignment to a concrete type")
	}
}

// ---- buffer helpers ----

func TestNewInputBuffer(t *testing.T) {
	buf := newInputBuffer("test data")
	if buf.String() != "test data" {
		t.Errorf("expected 'test data', got %q", buf.String())
	}
}

func TestNewOutputBuffer(t *testing.T) {
	buf := newOutputBuffer()
	if buf.Len() != 0 {
		t.Errorf("expected empty buffer, got %q", buf.String())
	}
}

// ---- buffer copy and truncation (standard library interfaces) ----

func TestBufferCopyAll(t *testing.T) {
	in := newInputBuffer("hello world")
	out := newOutputBuffer()

	_, err := io.Copy(out, in)
	if err != nil {
		t.Fatalf("io.Copy returned error: %v", err)
	}
	if out.String() != "hello world" {
		t.Errorf("expected 'hello world', got %q", out.String())
	}
}

func TestBufferCopyPartial(t *testing.T) {
	in := newInputBuffer("hello world")
	out := newOutputBuffer()

	_, err := io.CopyN(out, in, 5)
	if err != nil {
		t.Fatalf("io.CopyN returned error: %v", err)
	}
	if out.String() != "hello" {
		t.Errorf("expected 'hello', got %q", out.String())
	}
}

func TestBufferTruncate(t *testing.T) {
	in := newInputBuffer("hello world")
	out := newOutputBuffer()

	_, err := io.Copy(out, in)
	if err != nil {
		t.Fatalf("io.Copy returned error: %v", err)
	}
	out.Truncate(5)
	if out.String() != "hello" {
		t.Errorf("expected 'hello' after truncation, got %q", out.String())
	}
}
