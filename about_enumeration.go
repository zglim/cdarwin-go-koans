package go_koans

func aboutEnumeration() {
	{
		// Use the shared greetings list so the example data matches
		// the other collection koans.
		greetings := sampleGreetings

		var concatenated string
		var total int
		for i, v := range greetings {
			total += i
			concatenated += v
		}

		assert(concatenated == __string__) // for loops have a modern variation
		assert(total == __int__)           // which offers both a value and an index
	}

	{
		greetings := sampleGreetings

		var totalLen int
		for _, v := range greetings {
			totalLen += len(v)
		}

		assert(totalLen == __int__) // although we may omit either value
	}
}
