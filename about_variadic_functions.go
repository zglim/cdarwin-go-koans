package go_koans

import "strings"

// concatNames joins the given names with sep.  Variadic parameters are
// really just slices.
func concatNames(sep string, names ...string) string {
	return strings.Join(names, sep)
}

func aboutVariadicFunctions() {
	{
		// Pass each name as a separate argument.
		str := concatNames(" ", sampleNames[0], sampleNames[1], sampleNames[2])
		assert(str == __string__) // several values can be passed to variadic parameters
	}

	{
		// Unpack the shared slice with the spread operator.
		str := concatNames("-", sampleNames...)
		assert(str == __string__) // or a slice can be dotted in place of all of them
	}
}
