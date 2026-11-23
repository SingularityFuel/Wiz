package wiz

import (
	"github.com/SingularityFuel/wiz/string"
)

/*
	All hail the wiz module.

	Try to hide all the functionality in modules that export a type.
	Bonus points if user never has to import those modules individually.

	Meanwhile, this file holds constructors for the module types.
*/

// ASCII string
func String(x ...any) string.T {
	return string.N(x...)
}

// One-line ASCII string
func Line(x ...any) string.T {
	return string.L(x...)
}
