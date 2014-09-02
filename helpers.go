// Copyright 2014 Ryan Rogers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decimal

import "math"

func printedLength(n uint64) int {
	if n == 0 {
		return 1
	}
	// FIXME/NOTE: This is somewhat slow. Since we know that we're limited to
	// the range of a uint64, it could be replaced with a big (ugly) if tree.
	// I'm not sure the need for that actually exists at the moment.
	return int(math.Floor(math.Log10(float64(n))) + 1)
}

func simplifyNumber(n uint64) (uint64, int) {
	for n >= 10 && n%10 == 0 {
		n /= 10
	}
	return n, printedLength(n)
}
