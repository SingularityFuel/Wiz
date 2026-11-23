package string

/*
	Module type's exposed methods should have readable, usable signatures.
*/

// String's length in ascii characters.
func (t T) Length() (length int) {
	return len(t.txt)
}

// String's literal text as ascii characters.
// Satisfies fmt.Stringer interface.
func (t T) String() (ascii string) {
	return t.txt
}
