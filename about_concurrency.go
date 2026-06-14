package go_koans

// primeSearchCeiling bounds findPrimeNumbers so a runaway search can never climb
// forever. The koan jokes that "i is afraid of heights"; this constant is that
// fear made explicit, and naming it keeps the limit in one obvious place.
const primeSearchCeiling = 100

// isPrimeNumber reports whether possiblePrime has no divisors other than 1 and
// itself. It is pure arithmetic: no channels, no goroutines, just the math, so
// the "is it prime?" question stays separate from how primes are delivered.
func isPrimeNumber(possiblePrime int) bool {
	for underPrime := 2; underPrime < possiblePrime; underPrime++ {
		if possiblePrime%underPrime == 0 {
			return false
		}
	}
	return true
}

// guardSearchCeiling stops the producer from searching past primeSearchCeiling.
// Pulling the bound into its own named step makes "termination control" a
// distinct responsibility instead of a bare assert buried in the loop body,
// while keeping the original crash-if-too-high safety semantics intact.
func guardSearchCeiling(candidate int) {
	assert(candidate < primeSearchCeiling) // i is afraid of heights
}

// findPrimeNumbers is the producer goroutine: it walks the integers, sends each
// prime onto channel, and guards the ceiling on every step. Detection, sending,
// and the safety bound are each delegated, so the loop reads as three teaching
// steps and nothing more.
func findPrimeNumbers(channel chan int) {
	for candidate := 2; ; candidate++ {
		if isPrimeNumber(candidate) {
			channel <- candidate
		}
		guardSearchCeiling(candidate)
	}
}

func aboutConcurrency() {
	ch := make(chan int)

	go findPrimeNumbers(ch) // concurrency can be almost trivial

	assert(<-ch == 2)
	assert(<-ch == 3)
	assert(<-ch == 5)
	assert(<-ch == 7)
	assert(<-ch == 11)
}
