package go_koans

// person is the shared struct model for the value/pointer koans. It lives at
// package level (like coolNumber) so the structs and allocation koans use the
// very same type instead of redefining it locally.
type person struct {
	name string
	age  int
}

func aboutStructs() {
	var bob struct {
		name string
		age  int
	}
	bob.name = "bob"
	bob.age = 30

	assert(bob.name == __string__) // structs are collections of named variables
	assert(bob.age == __int__)     // each field has both setter and getter behavior

	var john person
	john.name = "bob"
	john.age = __int__

	assert(bob == john) // assuredly, bob is certainly not john.. yet
}
