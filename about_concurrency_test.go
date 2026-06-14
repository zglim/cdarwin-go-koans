package go_koans

import (
	"testing"
	"time"
)

// --- Buffered channel helpers ---

func TestSendStrings_Single(t *testing.T) {
	ch := make(chan string, 1)
	go sendStrings(ch, "foo")

	val := receiveString(ch)
	if val != "foo" {
		t.Errorf("expected 'foo', got '%s'", val)
	}
}

func TestSendStrings_Multiple(t *testing.T) {
	ch := make(chan string, 3)
	go sendStrings(ch, "a", "b", "c")

	// Allow goroutine to send all values
	time.Sleep(10 * time.Millisecond)

	if len(ch) != 3 {
		t.Errorf("expected 3 items in buffered channel, got %d", len(ch))
	}

	for _, expected := range []string{"a", "b", "c"} {
		val := receiveString(ch)
		if val != expected {
			t.Errorf("expected '%s', got '%s'", expected, val)
		}
	}
}

func TestReceiveString(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "hello"

	val := receiveString(ch)
	if val != "hello" {
		t.Errorf("expected 'hello', got '%s'", val)
	}
	if len(ch) != 0 {
		t.Errorf("expected empty channel after receive, got len %d", len(ch))
	}
}

func TestDrainChannel(t *testing.T) {
	ch := make(chan string, 3)
	ch <- "x"
	ch <- "y"
	ch <- "z"

	drainChannel(ch, 2)

	if len(ch) != 1 {
		t.Errorf("expected 1 item after draining 2, got %d", len(ch))
	}

	remaining := receiveString(ch)
	if remaining != "z" {
		t.Errorf("expected 'z' to remain, got '%s'", remaining)
	}
}

func TestBufferedChannel_LengthTransitions(t *testing.T) {
	ch := make(chan string, 2)

	// Empty buffered channel has length 0
	if len(ch) != 0 {
		t.Errorf("new buffered channel should have len 0, got %d", len(ch))
	}

	// After one send, length is 1
	ch <- "foo"
	if len(ch) != 1 {
		t.Errorf("expected len 1 after one send, got %d", len(ch))
	}

	// After receive, length drops back to 0
	val := <-ch
	if val != "foo" {
		t.Errorf("expected 'foo', got '%s'", val)
	}
	if len(ch) != 0 {
		t.Errorf("expected len 0 after receive, got %d", len(ch))
	}
}

func TestBufferedChannel_DrainMakesRoom(t *testing.T) {
	ch := make(chan string, 2)

	// Fill the buffer completely
	ch <- "bar"
	ch <- "quux"

	if len(ch) != 2 {
		t.Fatalf("expected 2 items, got %d", len(ch))
	}

	// Drain one slot in a goroutine so the next send can proceed
	done := make(chan struct{})
	go func() {
		drainChannel(ch, 1)
		close(done)
	}()

	// This send would deadlock without the drain above
	ch <- "extra"

	<-done

	if len(ch) != 2 {
		t.Errorf("expected 2 items after drain+send, got %d", len(ch))
	}
}

// --- Prime generation helpers ---

func TestIsPrimeNumber_Primes(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	for _, p := range primes {
		if !isPrimeNumber(p) {
			t.Errorf("expected %d to be prime", p)
		}
	}
}

func TestIsPrimeNumber_NonPrimes(t *testing.T) {
	nonPrimes := []int{4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20, 21}
	for _, n := range nonPrimes {
		if isPrimeNumber(n) {
			t.Errorf("expected %d to NOT be prime", n)
		}
	}
}

func TestSendPrimes_FirstFive(t *testing.T) {
	ch := make(chan int, 5)
	go sendPrimes(ch, 5, 100)

	expected := []int{2, 3, 5, 7, 11}
	primes := collectPrimes(ch)

	if len(primes) != len(expected) {
		t.Fatalf("expected %d primes, got %d: %v", len(expected), len(primes), primes)
	}

	for i, p := range primes {
		if p != expected[i] {
			t.Errorf("prime[%d]: expected %d, got %d", i, expected[i], p)
		}
	}
}

func TestSendPrimes_SinglePrime(t *testing.T) {
	ch := make(chan int, 1)
	go sendPrimes(ch, 1, 100)

	primes := collectPrimes(ch)
	if len(primes) != 1 || primes[0] != 2 {
		t.Errorf("expected [2], got %v", primes)
	}
}

func TestSendPrimes_ClosesChannel(t *testing.T) {
	ch := make(chan int, 10)
	go sendPrimes(ch, 3, 100)

	// collectPrimes blocks until channel is closed, so this tests
	// that sendPrimes properly closes ch after sending all primes.
	primes := collectPrimes(ch)
	if len(primes) != 3 {
		t.Errorf("expected 3 primes, got %d", len(primes))
	}
}

func TestSendPrimes_RespectsMaxCandidate(t *testing.T) {
	ch := make(chan int, 100)
	go sendPrimes(ch, 1000, 30) // ask for 1000 primes but cap candidate at 30

	primes := collectPrimes(ch)
	// Primes under 30: 2,3,5,7,11,13,17,19,23,29 = 10 primes
	expectedCount := 10
	if len(primes) != expectedCount {
		t.Errorf("expected %d primes (bounded by maxCandidate=30), got %d: %v",
			expectedCount, len(primes), primes)
	}

	// Verify the last prime is still correct
	if len(primes) > 0 && primes[len(primes)-1] != 29 {
		t.Errorf("expected last prime to be 29, got %d", primes[len(primes)-1])
	}
}

func TestSendPrimes_OrderPreserved(t *testing.T) {
	ch := make(chan int, 20)
	go sendPrimes(ch, 10, 100)

	primes := collectPrimes(ch)

	// Verify strictly increasing order
	for i := 1; i < len(primes); i++ {
		if primes[i] <= primes[i-1] {
			t.Errorf("primes not in increasing order at index %d: %d <= %d",
				i, primes[i], primes[i-1])
		}
	}
}

func TestCollectPrimes_EmptyChannel(t *testing.T) {
	ch := make(chan int)
	close(ch) // close immediately, no values sent

	primes := collectPrimes(ch)
	if len(primes) != 0 {
		t.Errorf("expected empty slice from closed empty channel, got %v", primes)
	}
}

// --- Goroutine-driven send path ---

func TestGoroutineDrivenSend_PrimeOrder(t *testing.T) {
	ch := make(chan int, 5)
	go sendPrimes(ch, 5, 100)

	// Receive primes one by one and verify order matches 2,3,5,7,11
	expected := []int{2, 3, 5, 7, 11}
	for i, want := range expected {
		got := <-ch
		if got != want {
			t.Errorf("prime[%d]: expected %d, got %d", i, want, got)
		}
	}
}

func TestGoroutineDrivenSend_WithFindPrimeNumbers(t *testing.T) {
	// Verify that findPrimeNumbers (once filled in by the student)
	// would produce the same sequence when using sendPrimes internally.
	// Here we test the sendPrimes path directly to validate the plumbing.
	ch := make(chan int)
	go sendPrimes(ch, 5, 100)

	results := make([]int, 0, 5)
	for p := range ch {
		results = append(results, p)
	}

	expected := []int{2, 3, 5, 7, 11}
	if len(results) != len(expected) {
		t.Fatalf("expected %d primes, got %d", len(expected), len(results))
	}
	for i, p := range results {
		if p != expected[i] {
			t.Errorf("result[%d]: expected %d, got %d", i, expected[i], p)
		}
	}
}
