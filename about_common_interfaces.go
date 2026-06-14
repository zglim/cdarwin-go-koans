package go_koans

import "bytes"

// ---- buffer helpers ----
//
// These helpers reduce the repeated "create input buffer, create output buffer"
// boilerplate that would otherwise appear in every common-interface koan block.

// newInputBuffer creates a *bytes.Buffer pre-loaded with the given content.
// It implements io.Reader, so it can be used as a data source.
func newInputBuffer(content string) *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.WriteString(content)
	return buf
}

// newOutputBuffer creates an empty *bytes.Buffer ready to receive writes.
// It implements io.Writer, so it can be used as a data sink.
func newOutputBuffer() *bytes.Buffer {
	return new(bytes.Buffer)
}

// ---- koan: standard library interfaces (io.Reader / io.Writer) ----
//
// bytes.Buffer implements both io.Reader and io.Writer.
// Your task: use the standard library to move data from the input buffer
// to the output buffer. Think about which io package functions fit each case.

func aboutCommonInterfaces() {
	{
		// Copy all data from an io.Reader to an io.Writer.
		in := newInputBuffer("hello world")
		out := newOutputBuffer()

		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		_ = in                              // use 'in' as your io.Reader source
		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		// Copy only a portion of data from an io.Reader to an io.Writer.
		in := newInputBuffer("hello world")
		out := newOutputBuffer()

		_ = in                           // use 'in' as your io.Reader source
		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
