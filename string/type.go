package string

import "fmt"

/*
	Module type is always T.
	Constructors are single letters.
	Readability is not required, as user should be importing wiz.
*/

// String type. Allows ASCII only.
type T struct {
	txt string
}

// Interface check
var _ fmt.Stringer = Z()

// Zero-value string is empty.
func Z() T {
	return T{}
}

// New string from anything. Checks for literal or Stringer interface,
// but if not, just lists object's type.
// Replaces non-printable ascii with ?. Can contain tabs and newlines.
// Multiple inputs will be separated with spaces. Empty inputs will be skipped in such lists.
func N(x ...any) T {
	return n(spacem(ss(x)...))
}

// New single line string from anything. Checks for literal or Stringer interface,
// but if not, just lists object's type.
// Replaces non-printable ascii with ?. Replaces newlines and tabs with single spaces.
// Multiple inputs will be separated with spaces. Empty inputs will be skipped in such lists.
func L(x ...any) T {
	return l(spacem(ss(x)...))
}
