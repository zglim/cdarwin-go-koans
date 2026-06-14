package go_koans

func aboutInterfaces() {
	bob := new(human)     // bob is a kind of *human
	rspec := new(program) // rspec is a kind of *program

	assert(runner(bob) == __runner__) // conformed interfaces need not be declared, they are inferred

	assert(bob.milesCompleted == 0)
	assert(rspec.executionCount == 0)

	runTwice(bob)   // bob fits the profile for a 'runner'
	runTwice(rspec) // rspec also fits the profile for a 'runner'

	assert(bob.milesCompleted == __int__)   // bob is affected by running in his own unique way (probably fatigue)
	assert(rspec.executionCount == __int__) // rspec can run completely differently than bob, thanks to interfaces
}

// abstract interface and the shared driver that requires it.
//
// runTwice() is to the runner koans what copyDemo() (in
// about_common_interfaces.go) is to the io.Reader / io.Writer koans: it takes a
// value behind the abstract interface and triggers its behavior, without caring
// which concrete type is actually hiding underneath.

type runner interface {
	run()
}

func runTwice(r runner) {
	r.run()
	r.run()
}

// concrete type implementing the interface

type human struct {
	milesCompleted int
}

func (h *human) run() {
	h.milesCompleted++
}

// another concrete type implementing the interface

type program struct {
	executionCount int
}

func (p *program) run() {
	p.executionCount++
}
