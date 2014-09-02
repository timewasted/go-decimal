// Copyright 2014 Ryan Rogers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package decimal allows for working with decimal values of almost any
// precision.
package decimal

import (
	"fmt"
	"math"
)

// Bounds checking values.
const (
	minSignedInt64   = -(1 << 63)
	maxSignedInt64   = 1<<63 - 1
	maxUnsignedInt64 = 1<<64 - 1
)

// DecimalSeparator is the character to use for a decimal separator.
var DecimalSeparator = '.'

// ThousandsSeparator is the character to use for a thousands separator.
var ThousandsSeparator = ','

// Decimal is a representation of a Decimal value.
type Decimal struct {
	Valid, Negative        bool
	numerator, denominator uint64
	denominatorDigits      int
}

// ParseDecimal converts the string s into a Decimal. A valid Decimal string
// has the following format:
//
// SNN.DD
//
// S is a negative (-) or positive (+) sign (optional)
// NN is zero or more decimal digits (up to the max value for a uint64)
// . is the defined DecimalSeparator (default .)
// DD is zero or more decimal digits (up to the max value for a uint64)
//
// NN or DD can be omitted, but not both.
func ParseDecimal(s string) (*Decimal, error) {
	const fnName = "ParseDecimal"

	if len(s) == 0 {
		return nil, syntaxError(fnName, s)
	}

	decimal := &Decimal{}

	i := 0
	if s[0] == '+' {
		i = 1
	} else if s[0] == '-' {
		i = 1
		decimal.Negative = true
	}

	denominatorDigits := -1
	for ; i < len(s); i++ {
		var v uint8
		d := s[i]
		switch {
		case '0' <= d && d <= '9':
			v = uint8(d - '0')
		case d == uint8(DecimalSeparator):
			if denominatorDigits != -1 {
				return nil, syntaxError(fnName, s)
			}
			denominatorDigits = 0
			continue
		default:
			return nil, syntaxError(fnName, s)
		}

		// Bounds checking.
		var n *uint64
		if denominatorDigits == -1 {
			n = &decimal.numerator
		} else {
			n = &decimal.denominator
			denominatorDigits++
		}
		newN := *n*10 + uint64(v)
		// FIXME: I'm pretty sure newN can never be > maxUnsignedInt64.
		if newN < *n || newN > maxUnsignedInt64 {
			return nil, rangeError(fnName, s)
		}
		*n = newN
		decimal.Valid = true
	}

	if decimal.Valid {
		if denominatorDigits != -1 {
			decimal.denominatorDigits = denominatorDigits
		}
		return decimal, nil
	}
	return nil, syntaxError(fnName, s)
}

// Cmp compares d1 and d2 and returns:
//
//   -1 if d1 <  d2
//    0 if d1 == d2
//   +1 if d1 >  d2
//
func (d1 *Decimal) Cmp(d2 *Decimal) (r int) {
	if d1.Negative == d2.Negative {
		if d1.numerator == d2.numerator && d1.denominator == d2.denominator {
			return
		}
		if d1.numerator > d2.numerator || d1.numerator == d2.numerator && d1.denominator > d2.denominator {
			r = 1
		} else {
			r = -1
		}

		if d1.Negative {
			r = -r
		}
	} else {
		r = 1
		if d1.Negative {
			r = -r
		}
	}
	return
}

// Add sets d1 to the sum of d1+d2. An error is returned if either d1 or d2
// are flagged as being invalid, or if the operation would result in d1
// overflowing. d1 is unchanged on error.
func (d1 *Decimal) Add(d2 *Decimal) error {
	if !d1.Valid || !d2.Valid {
		return ErrNotValid
	}

	// Bounds checking.
	// FIXME: I'm pretty sure newV can never be > maxUnsignedInt64.
	newV := d1.denominator + d2.denominator
	if newV < d1.denominator || newV > maxUnsignedInt64 {
		return rangeError("Add", d1.String()+" + "+d2.String())
	}
	newV = d1.numerator + d2.numerator
	if newV < d1.numerator || newV > maxUnsignedInt64 {
		return rangeError("Add", d1.String()+" + "+d2.String())
	}

	// Ensure equal "length" denominators.
	d2Denominator := d2.denominator
	if d1.denominatorDigits > d2.denominatorDigits {
		d2Denominator *= uint64(math.Pow10(d1.denominatorDigits - d2.denominatorDigits))
	} else if d2.denominatorDigits > d1.denominatorDigits {
		d1.denominator *= uint64(math.Pow10(d2.denominatorDigits - d1.denominatorDigits))
		d1.denominatorDigits = d2.denominatorDigits
	}

	if d1.Negative == d2.Negative {
		d1.denominator += d2Denominator
		d1.numerator += d2.numerator
	} else {
		n1, n2 := d1.Negative, d2.Negative
		d1.Negative, d2.Negative = false, false
		if d1.Cmp(d2) >= 0 {
			d1.denominator -= d2Denominator
			d1.numerator -= d2.numerator
		} else {
			d1.denominator = d2Denominator - d1.denominator
			d1.numerator = d2.numerator - d1.numerator
			n1 = !n1
		}
		d1.Negative, d2.Negative = n1, n2
	}

	// Perform a carry, if needed.
	d1DigitsNew := printedLength(d1.denominator)
	if d1DigitsNew > d2.denominatorDigits {
		mod := uint64(math.Pow10(d1.denominatorDigits))
		d1.numerator += d1.denominator / mod
		d1.denominator %= mod
	}

	// Simplify d1 after performing the carry.
	d1.denominator, d1DigitsNew = simplifyNumber(d1.denominator)
	d1.denominatorDigits = d1DigitsNew
	return nil
}

// String returns the string representation of the Decimal. Thousands
// separators are not used.
func (d *Decimal) String() string {
	const fmtString = "%%d%%c%%0%dd"
	if d.Negative {
		return fmt.Sprintf("-"+fmt.Sprintf(fmtString, d.denominatorDigits), d.numerator, DecimalSeparator, d.denominator)
	}
	return fmt.Sprintf(fmt.Sprintf(fmtString, d.denominatorDigits), d.numerator, DecimalSeparator, d.denominator)
}

// FormattedString returns the string representation of the Decimal. Thousands
// separators are used.
func (d *Decimal) FormattedString() string {
	if d.numerator < 1000 {
		return d.String()
	}

	numerator := fmt.Sprintf("%d", d.numerator)
	var pn []byte
	if len(numerator)%3 != 0 {
		pn = make([]byte, len(numerator)+len(numerator)/3)
	} else {
		pn = make([]byte, len(numerator)+len(numerator)/3-1)
	}
	pnIdx := 0

	start := 0
	for i := len(numerator) % 3; i <= len(numerator); i += 3 {
		if i == 0 {
			continue
		}
		pnIdx += copy(pn[pnIdx:], numerator[start:i])
		if i != len(numerator) {
			pnIdx += copy(pn[pnIdx:], string(ThousandsSeparator))
		}
		start = i
	}

	const fmtString = "%%s%%c%%0%dd"
	if d.Negative {
		return fmt.Sprintf("-"+fmt.Sprintf(fmtString, d.denominatorDigits), string(pn), DecimalSeparator, d.denominator)
	}
	return fmt.Sprintf(fmt.Sprintf(fmtString, d.denominatorDigits), string(pn), DecimalSeparator, d.denominator)
}
