// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = /*not a triangle*/ iota
	Equ /*equilateral*/
	Iso /*isosceles*/
	Sca /*scalene*/
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var k Kind

	if a == 0 || b == 0 || c == 0 || a+b < c || a+c < b || b+c < a {
		// Not triangle
		k = 0
	} else if a == b && b == c {
		// equilateral
		k = 1
	} else if a == b || b == c || c == a {
		// isosceles
		k = 2
	} else if a != b && a != c && b != a {
		// scalene
		k = 3
	}

	return k
}
