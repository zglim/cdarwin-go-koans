package go_koans

import (
	"bytes"
	"io"
)

func aboutCommonInterfaces() {
	// Just like a *human or a *program satisfies our own runner interface over
	// in about_interfaces.go, the standard library's *bytes.Buffer satisfies
	// io.Reader and io.Writer. copyDemo() puts that concrete *bytes.Buffer
	// behind those abstract interfaces, runs the copy operation you supply, and
	// reports whatever ended up in the destination -- the same "concrete type
	// behind an abstract interface, then trigger behavior" idea as runTwice().

	assert(copyDemo("hello world", func(dst io.Writer, src io.Reader) {
		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/
	}) == "hello world") // get data from the io.Reader to the io.Writer

	assert(copyDemo("hello world", func(dst io.Writer, src io.Reader) {
		// Your code goes here.
	}) == "hello") // duplicate only a portion of the io.Reader
}

// copyDemo is the shared driver for the io koans above. It prepares the
// input/output buffers once, hands them to the supplied copy operation as an
// io.Reader / io.Writer pair, and returns whatever landed in the destination.
// It plays the same role here that runTwice() plays for the runner koans.
func copyDemo(input string, copyData func(dst io.Writer, src io.Reader)) string {
	in := bytes.NewBufferString(input) // in is an io.Reader already holding our source bytes
	out := new(bytes.Buffer)           // out is an io.Writer we copy into
	copyData(out, in)
	return out.String()
}
