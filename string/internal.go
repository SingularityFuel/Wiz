package string

import (
	"fmt"
	"strings"

	u "github.com/SingularityFuel/wiz/u"
)

// internal converter
func s(x any) string {
	if u.Nil(x) {
		return "nil"
	}
	if t, ok := x.(string); ok {
		return t
	}
	if t, ok := x.(fmt.Stringer); ok {
		return t.String()
	}
	return u.Type(x)
}

// internal list converter
func ss(x ...any) []string {
	o := make([]string, len(x))
	for i, y := range x {
		o[i] = s(y)
	}
	return o
}

// internal string combiner. skips empty. separates with space.
func spacem(x ...string) string {
	o := ""
	for _, y := range x {
		if y != "" {
			if o != "" {
				o += " "
			}
			o += y
		}
	}
	return o
}

// internal constructor
func n(s string) T {
	return T{txt: qascii(s, false, false)}
}

// internal constructor
func l(s string) T {
	return T{txt: qascii(s, true, true)}
}

// internal ascii|? replacer which may also convert tabs and newlines to single spaces
func qascii(s string, replaceTabs, replaceLines bool) string {
	var b strings.Builder
	b.Grow(len(s)) // pre-allocate to avoid reallocations
	for _, r := range s {
		switch {
		case r == '\n': // newline
			if replaceLines {
				b.WriteRune(' ')
			} else {
				b.WriteRune(r)
			}
		case r == '\t': // tab
			if replaceTabs {
				b.WriteRune(' ')
			} else {
				b.WriteRune(r)
			}
		case r >= 0x20 && r <= 0x7E: // printable ASCII (including space)
			b.WriteRune(r)
		default:
			b.WriteByte('?') // everything else becomes ?
		}
	}
	return b.String()
}
