package string

import (
	"fmt"
	"strings"

	u "github.com/SingularityFuel/wiz/u"
)

// String type. Let's allow ASCII only.
type T struct {
	txt string
}

// Zero-value string is empty.
func Z() T {
	return T{}
}

// New string from anything. Checks for literal or Stringer interface, but if not, just lists object's type.
// Replaces non-ascii with ?. Can contain tabs and newlines.
func N(x any) T {
	if u.Nil(x) {
		return qascii("nil")
	}
	if s, ok := x.(string); ok {
		return qascii(s)
	}
	if s, ok := x.(fmt.Stringer); ok {
		return qascii(s.String())
	}
	return qascii(u.Type(x))
}

func qascii(s string) T {
	var b strings.Builder
	b.Grow(len(s)) // pre-allocate to avoid reallocations
	for _, r := range s {
		switch {
		case r == '\n' || r == '\t': // newline, tab
			b.WriteRune(r)
		case r >= 0x20 && r <= 0x7E: // printable ASCII (including space)
			b.WriteRune(r)
		default:
			b.WriteByte('?') // everything else becomes ?
		}
	}
	return T{txt: b.String()}
}
