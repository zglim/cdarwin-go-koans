package go_koans

// ---- abstract interface and utility function ----

// runner is satisfied by any type that has a run() method.
// No explicit "implements" declaration is needed — Go infers it.
type runner interface {
	run()
}

// runTwice demonstrates polymorphism: it works on any runner,
// regardless of the concrete type behind the interface.
func runTwice(r runner) {
	r.run()
	r.run()
}

// ---- concrete implementations ----

// human implements runner implicitly via its run() method.
type human struct {
	milesCompleted int
}

func (h *human) run() {
	h.milesCompleted++
}

// program also implements runner implicitly, but with completely different behavior.
type program struct {
	executionCount int
}

func (p *program) run() {
	p.executionCount++
}

// ---- koan: implicit interface implementation ----

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
