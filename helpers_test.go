// Copyright 2014 Ryan Rogers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decimal

import "testing"

func TestSimplifyNumber(t *testing.T) {
	// NOTE: This also tests printedLength().

	type testResult struct {
		number uint64
		digits int
	}
	tests := map[uint64]testResult{
		0: testResult{
			number: 0,
			digits: 1,
		},
		1: testResult{
			number: 1,
			digits: 1,
		},
		11: testResult{
			number: 11,
			digits: 2,
		},
		111: testResult{
			number: 111,
			digits: 3,
		},
		1111: testResult{
			number: 1111,
			digits: 4,
		},
		11111: testResult{
			number: 11111,
			digits: 5,
		},
		111111: testResult{
			number: 111111,
			digits: 6,
		},
		1111111: testResult{
			number: 1111111,
			digits: 7,
		},
		11111111: testResult{
			number: 11111111,
			digits: 8,
		},
		111111111: testResult{
			number: 111111111,
			digits: 9,
		},
		1111111111: testResult{
			number: 1111111111,
			digits: 10,
		},
		11111111111: testResult{
			number: 11111111111,
			digits: 11,
		},
		111111111111: testResult{
			number: 111111111111,
			digits: 12,
		},
		1111111111111: testResult{
			number: 1111111111111,
			digits: 13,
		},
		11111111111111: testResult{
			number: 11111111111111,
			digits: 14,
		},
		111111111111111: testResult{
			number: 111111111111111,
			digits: 15,
		},
		1111111111111111: testResult{
			number: 1111111111111111,
			digits: 16,
		},
		11111111111111111: testResult{
			number: 11111111111111111,
			digits: 17,
		},
		111111111111111111: testResult{
			number: 111111111111111111,
			digits: 18,
		},
		1111111111111111111: testResult{
			number: 1111111111111111111,
			digits: 19,
		},
		11111111111111111111: testResult{
			number: 11111111111111111111,
			digits: 20,
		},
		10000000000000000000: testResult{
			number: 1,
			digits: 1,
		},
		10000000000000000001: testResult{
			number: 10000000000000000001,
			digits: 20,
		},
	}

	for value, result := range tests {
		n, d := simplifyNumber(value)
		if result.number != n {
			t.Errorf("Expected %d to return %d, received %d.", value, result.number, n)
		}
		if result.digits != d {
			t.Errorf("Expected %d to return %d, received %d.", value, result.digits, d)
		}
	}
}
