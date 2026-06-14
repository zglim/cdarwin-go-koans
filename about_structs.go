package go_koans

func aboutStructs() {
	// bob is created using the shared person type (see shared_types.go).
	var bob person
	bob.name = "bob"
	bob.age = 30

	assert(bob.name == __string__) // structs are collections of named variables
	assert(bob.age == __int__)     // each field has both setter and getter behavior

	// john uses the same person type – no need to redefine it locally.
	var john person
	john.name = "bob"
	john.age = __int__

	assert(bob == john) // assuredly, bob is certainly not john.. yet
}
