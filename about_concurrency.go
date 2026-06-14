package go_koans

// isPrimeNumber reports whether n is a prime number.
func isPrimeNumber(n int) bool {
	for d := 2; d < n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

// sendPrimes sends exactly count prime numbers to ch, starting from 2.
// The maxCandidate bound prevents runaway computation (height safety).
// When done, it closes ch to signal that no more values will be sent.
func sendPrimes(ch chan int, count int, maxCandidate int) {
	sent := 0
	for candidate := 2; sent < count && candidate < maxCandidate; candidate++ {
		if isPrimeNumber(candidate) {
			ch <- candidate
			sent++
		}
	}
	close(ch)
}

// collectPrimes reads all values from ch until the channel is closed,
// returning them in receive order.
func collectPrimes(ch chan int) []int {
	var primes []int
	for p := range ch {
		primes = append(primes, p)
	}
	return primes
}

// findPrimeNumbers sends prime numbers to ch using the sendPrimes helper.
func findPrimeNumbers(ch chan int) {
	// your code goes here
	// hint: call sendPrimes(ch, 5, 100) to send the first 5 primes
}

func aboutConcurrency() {
	ch := make(chan int)

	assert(__delete_me__) // concurrency can be almost trivial
	// your code goes here
	// hint: launch findPrimeNumbers(ch) in a goroutine with "go findPrimeNumbers(ch)"

	assert(<-ch == 2)
	assert(<-ch == 3)
	assert(<-ch == 5)
	assert(<-ch == 7)
	assert(<-ch == 11)
}
