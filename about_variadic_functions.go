package go_koans

import "strings"

func concatNames(sep string, names ...string) string {
	return strings.Join(names, sep) // variadic parameters are really just slices
}

func aboutVariadicFunctions() {
	{
		friends := friendNames()
		str := concatNames(" ", friends[0], friends[1], friends[2])
		assert(str == __string__) // several values can be passed to variadic parameters
	}

	{
		friends := friendNames()
		str := concatNames("-", friends...)
		assert(str == __string__) // or a slice can be dotted in place of all of them
	}
}
