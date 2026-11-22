package flow

import "time"

/*
	State
*/

type Flow struct {
	state *flow // Nil is invalid. Only properly initialized flows should be usable.
}

type flow struct {
	cancelled bool       // If cancelled without error, this is true.
	err       string     // If aborted due to error, message is here. Can be "Deadline Exceeded".
	deadline  *time.Time // Flow self-cancels after deadline. Nil if no deadline.
}

/*
	Constructors
*/

func New() Flow {
	return Flow{
		state: &flow{},
	}
}

func WithDeadline(t time.Time) Flow {
	f := New()
	f.state.deadline = &t
	f.update()
	return f
}

/*
	Functionality
*/

// Returns true if flow should keep executing.
func (f Flow) Ok() bool {
	f.update()
	return f.ok()
}

// If flow is still ok, cancel it without error.
func (f Flow) Cancel() {
	f.update()
	if f.ok() {
		f.state.cancelled = true
	}
}

// If flow is still ok, stop it and remember error message.
func (f Flow) Fail(err string) {
	f.update()
	if f.ok() {
		f.state.err = err
	}
}

// If flow is still ok, stop it with error if any of the conditions are true.
func (f Flow) Forbid(err string, conditions ...bool) {
	for _, c := range conditions {
		if c {
			f.Fail("Forbid: " + err)
			return
		}
	}
}

// If flow is still ok, stop it with error if any of the conditions are false.
func (f Flow) Require(err string, conditions ...bool) {
	for _, c := range conditions {
		if !c {
			f.Fail("Require: " + err)
			return
		}
	}
}

/*
	Internal
*/

// Checks if deadline exceeded and updates state if so.
func (f Flow) update() {
	if f.ok() && f.state.deadline != nil && time.Now().After(*f.state.deadline) {
		f.state.err = "Deadline exceeded."
	}
}

// Field by field ok check. Unlike Ok(), does not check deadline against time.Now().
func (f Flow) ok() bool {
	if f.state == nil || f.state.cancelled || f.state.err != "" {
		return false
	}
	return true
}
