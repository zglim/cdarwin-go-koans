package go_koans

// makeRoomInQueue receives a single value from a buffered channel so that a
// later send has somewhere to land. Giving the goroutine body a name (instead
// of an inline anonymous function) keeps aboutChannels focused on *what*
// happens to the queue, while this helper owns *how* room is made without
// deadlocking. It mirrors findPrimeNumbers in about_concurrency.go: a small,
// named function that cooperates with the koan over a channel.
func makeRoomInQueue(queue chan string) {
	<-queue
}

func aboutChannels() {
	ch := make(chan string, 2)

	assert(len(ch) == 0) // channels are like buffers

	ch <- "foo" // i mean, "metaphors are like similes"

	assert(len(ch) == 1) // they can be queried for queued items

	assert(<-ch == "foo") // items can be popped out of them

	assert(len(ch) == 0) // and len() always reflects the "current" queue status

	// the 'go' keyword runs a function-call in a new "goroutine" which executes
	// "concurrently" with the calling "goroutine". Here that goroutine drains
	// one item so the third send below has room and we don't deadlock.
	go makeRoomInQueue(ch)

	ch <- "bar"   // this send will succeed
	ch <- "quux"  // there's enough room for this send too
	ch <- "extra" // the buffer only has two slots, so the goroutine must drain one
}
