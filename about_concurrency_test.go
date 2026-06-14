package go_koans

import (
	"testing"
	"time"
)

// testTimeout bounds the concurrent tests below so a regression that introduces
// a deadlock is reported as a failure instead of hanging the suite forever.
const testTimeout = time.Second

// TestIsPrimeNumber pins down the pure prime predicate that the producer relies
// on, so a change to the math is caught independently of the goroutine plumbing.
func TestIsPrimeNumber(t *testing.T) {
	cases := map[int]bool{
		2:  true,
		3:  true,
		4:  false,
		5:  true,
		7:  true,
		9:  false,
		11: true,
		12: false,
		13: true,
		15: false,
	}
	for n, want := range cases {
		if got := isPrimeNumber(n); got != want {
			t.Errorf("isPrimeNumber(%d) = %v, want %v", n, got, want)
		}
	}
}

// TestFindPrimeNumbersProducesOrderedPrimes drives the producer goroutine and
// asserts the prime sequence the koan teaches (2, 3, 5, 7, 11) arrives in order.
func TestFindPrimeNumbersProducesOrderedPrimes(t *testing.T) {
	ch := make(chan int)
	go findPrimeNumbers(ch)

	want := []int{2, 3, 5, 7, 11}
	for _, expected := range want {
		select {
		case got := <-ch:
			if got != expected {
				t.Fatalf("expected prime %d, got %d", expected, got)
			}
		case <-time.After(testTimeout):
			t.Fatalf("timed out waiting for prime %d", expected)
		}
	}
}

// TestBufferedChannelLengthAndMakeRoom mirrors aboutChannels: it checks the
// buffered queue's length transitions (0 -> 1 -> 0) and then proves that the
// goroutine-driven makeRoomInQueue drains a slot so three sends fit through a
// two-slot buffer without deadlocking.
func TestBufferedChannelLengthAndMakeRoom(t *testing.T) {
	ch := make(chan string, 2)

	if got := len(ch); got != 0 {
		t.Fatalf("new buffered channel: len = %d, want 0", got)
	}

	ch <- "foo"
	if got := len(ch); got != 1 {
		t.Fatalf("after one send: len = %d, want 1", got)
	}

	if got := <-ch; got != "foo" {
		t.Fatalf("popped %q, want %q", got, "foo")
	}
	if got := len(ch); got != 0 {
		t.Fatalf("after pop: len = %d, want 0", got)
	}

	// The goroutine must free one slot so the third send has somewhere to go.
	go makeRoomInQueue(ch)

	done := make(chan struct{})
	go func() {
		ch <- "bar"
		ch <- "quux"
		ch <- "extra"
		close(done)
	}()

	select {
	case <-done:
		// success: all three sends completed because a slot was drained.
	case <-time.After(testTimeout):
		t.Fatal("third send deadlocked: makeRoomInQueue did not drain the buffer")
	}
}

// TestMakeRoomInQueueDrainsOneValue isolates the goroutine send/receive path and
// confirms makeRoomInQueue removes exactly one queued value, leaving the rest.
func TestMakeRoomInQueueDrainsOneValue(t *testing.T) {
	ch := make(chan string, 2)
	ch <- "first"
	ch <- "second"

	drained := make(chan struct{})
	go func() {
		makeRoomInQueue(ch)
		close(drained)
	}()

	select {
	case <-drained:
	case <-time.After(testTimeout):
		t.Fatal("makeRoomInQueue did not receive from the channel")
	}

	if got := len(ch); got != 1 {
		t.Fatalf("after draining one value: len = %d, want 1", got)
	}
	if got := <-ch; got != "second" {
		t.Fatalf("makeRoomInQueue drained the wrong value; remaining = %q, want %q", got, "second")
	}
}
