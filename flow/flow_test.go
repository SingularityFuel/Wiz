package flow_test

import (
	"testing"
	"time"

	"github.com/SingularityFuel/Wiz/flow"
)

func TestFlowInternals(t *testing.T) {
	// nil (invalid)
	nf := flow.Flow{}
	// future deadline
	df := flow.WithDeadline(time.Now().AddDate(0, 0, 1))
	// past deadline
	pf := flow.WithDeadline(time.Now().AddDate(0, 0, -1))
	// ok, no deadline
	zf := flow.New()
	// stopped
	sf := flow.New()
	sf.Cancel()
	// errored
	ef := flow.New()
	ef.Fail("custom error")

	if nf.Ok() || pf.Ok() || sf.Ok() || ef.Ok() {
		t.FailNow()
	}
	if !df.Ok() || !zf.Ok() {
		t.FailNow()
	}
}
