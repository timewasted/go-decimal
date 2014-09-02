// This file's contents were more or less ripped straight out of the "strconv"
// package.

package decimal

import "errors"

// ErrRange indicates that a value is out of range for the target type.
var ErrRange = errors.New("value out of range")

// ErrSyntax indicates that a value does not have the right syntax for the
// target type.
var ErrSyntax = errors.New("invalid syntax")

// ErrNotValid indicates that a value has Valid set to false.
var ErrNotValid = errors.New("value is not valid")

// NumError records a failed conversion.
type NumError struct {
	Func string // the failing function
	Num  string // the input
	Err  error  // the reason the conversion failed
}

func (e *NumError) Error() string {
	return "decimal." + e.Func + ": parsing '" + e.Num + "': " + e.Err.Error()
}

func syntaxError(fn, str string) *NumError {
	return &NumError{fn, str, ErrSyntax}
}

func rangeError(fn, str string) *NumError {
	return &NumError{fn, str, ErrRange}
}
