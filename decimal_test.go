// Copyright 2014 Ryan Rogers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decimal

import "testing"

func TestParseDecimal(t *testing.T) {
	type testResult struct {
		shouldFail, negative bool
		output               string
	}
	type parseDecimalTest struct {
		description, input string
		result             testResult
	}

	tests := []parseDecimalTest{
		{
			description: "Empty string",
			input:       "",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Just a decimal point",
			input:       ".",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Negative zero",
			input:       "-0.0",
			result: testResult{
				output: "0.0",
			},
		},
		{
			description: "No decimal point or denominator (positive value)",
			input:       "123",
			result: testResult{
				output: "123.0",
			},
		},
		{
			description: "No decimal point or denominator (negative value)",
			input:       "-123",
			result: testResult{
				negative: true,
				output:   "-123.0",
			},
		},
		{
			// FIXME: Should this fail instead?
			description: "Decimal point but no denominator (positive value)",
			input:       "123.",
			result: testResult{
				output: "123.0",
			},
		},
		{
			// FIXME: Should this fail instead?
			description: "Decimal point but no denominator (negative value)",
			input:       "-123.",
			result: testResult{
				negative: true,
				output:   "-123.0",
			},
		},
		{
			description: "Single digit denominator (positive value)",
			input:       "123.4",
			result: testResult{
				output: "123.4",
			},
		},
		{
			description: "Single digit denominator (negative value)",
			input:       "-123.4",
			result: testResult{
				negative: true,
				output:   "-123.4",
			},
		},
		{
			description: "Double digit denominator (positive value)",
			input:       "123.45",
			result: testResult{
				output: "123.45",
			},
		},
		{
			description: "Double digit denominator (negative value)",
			input:       "-123.45",
			result: testResult{
				negative: true,
				output:   "-123.45",
			},
		},
		{
			description: "Multiple digit denominator (positive value)",
			input:       "123.4567",
			result: testResult{
				output: "123.4567",
			},
		},
		{
			description: "Multiple digit denominator (negative value)",
			input:       "-123.4567",
			result: testResult{
				negative: true,
				output:   "-123.4567",
			},
		},
		{
			description: "No numerator (positive value)",
			input:       ".123",
			result: testResult{
				output: "0.123",
			},
		},
		{
			description: "No numerator (negative value)",
			input:       "-.123",
			result: testResult{
				negative: true,
				output:   "-0.123",
			},
		},
		{
			description: "Negative denominator",
			input:       ".-123",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Negative denominator",
			input:       "123.-456",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Negative denominator",
			input:       "-123.-456",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Bounds checking numerator (positive value)",
			input:       "18446744073709551615",
			result: testResult{
				output: "18446744073709551615.0",
			},
		},
		{
			description: "Bounds checking numerator (negative value)",
			input:       "-18446744073709551615",
			result: testResult{
				negative: true,
				output:   "-18446744073709551615.0",
			},
		},
		{
			description: "Bounds checking numerator (positive value)",
			input:       "18446744073709551616",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Bounds checking numerator (negative value)",
			input:       "-18446744073709551616",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Bounds checking denominator (positive value)",
			input:       ".18446744073709551615",
			result: testResult{
				output: "0.18446744073709551615",
			},
		},
		{
			description: "Bounds checking denominator (negative value)",
			input:       "-.18446744073709551615",
			result: testResult{
				negative: true,
				output:   "-0.18446744073709551615",
			},
		},
		{
			description: "Bounds checking denominator (positive value)",
			input:       ".18446744073709551616",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Bounds checking denominator (negative value)",
			input:       "-.18446744073709551616",
			result: testResult{
				shouldFail: true,
			},
		},
	}

	for _, test := range tests {
		d, err := ParseDecimal(test.input)
		if err != nil {
			if !test.result.shouldFail {
				t.Errorf("%s (input '%s'): expected success, received error '%v'.", test.description, test.input, err)
			}
			continue
		}
		if test.result.shouldFail {
			t.Errorf("%s (input '%s'): expected failure.", test.description, test.input)
			continue
		}
		if !d.Valid {
			t.Errorf("%s (input '%s'): expected valid result.", test.description, test.input)
			continue
		}
		if test.result.negative && !d.Negative {
			t.Errorf("%s (input '%s'): expected negative value.", test.description, test.input)
		} else if !test.result.negative && d.Negative {
			t.Errorf("%s (input '%s'): expected positive value.", test.description, test.input)
		}
		if test.result.output != d.String() {
			t.Errorf("%s (input '%s'): expected '%s', received '%s'.", test.description, test.input, test.result.output, d.String())
		}
	}
}

func TestCmp(t *testing.T) {
	type cmpTest struct {
		description, input1, input2 string
		result                      int
	}
	// The following just makes the tests more readable.
	const (
		lessThan    = -1
		equalTo     = 0
		greaterThan = 1
	)
	resultString := map[int]string{
		lessThan:    "less than",
		equalTo:     "equal to",
		greaterThan: "greater than",
	}

	// FIXME: There are a lot of tests that are essentially testing nothing.
	// I'm opting to keep them in for now.
	tests := []cmpTest{
		{
			description: "Same sign, same numerator, same denominator",
			input1:      "111.222",
			input2:      "111.222",
			result:      equalTo,
		},
		{
			description: "Same sign, same numerator, same denominator",
			input1:      "-111.222",
			input2:      "-111.222",
			result:      equalTo,
		},
		{
			description: "Same sign, same numerator, larger denominator",
			input1:      "111.333",
			input2:      "111.222",
			result:      greaterThan,
		},
		{
			description: "Same sign, same numerator, larger denominator",
			input1:      "-111.333",
			input2:      "-111.222",
			result:      lessThan,
		},
		{
			description: "Same sign, same numerator, smaller denominator",
			input1:      "111.222",
			input2:      "111.333",
			result:      lessThan,
		},
		{
			description: "Same sign, same numerator, smaller denominator",
			input1:      "-111.222",
			input2:      "-111.333",
			result:      greaterThan,
		},
		{
			description: "Same sign, larger numerator, same denominator",
			input1:      "222.222",
			input2:      "111.222",
			result:      greaterThan,
		},
		{
			description: "Same sign, larger numerator, same denominator",
			input1:      "-222.222",
			input2:      "-111.222",
			result:      lessThan,
		},
		{
			description: "Same sign, larger numerator, larger denominator",
			input1:      "222.333",
			input2:      "111.222",
			result:      greaterThan,
		},
		{
			description: "Same sign, larger numerator, larger denominator",
			input1:      "-222.333",
			input2:      "-111.222",
			result:      lessThan,
		},
		{
			description: "Same sign, larger numerator, smaller denominator",
			input1:      "222.111",
			input2:      "111.222",
			result:      greaterThan,
		},
		{
			description: "Same sign, larger numerator, smaller denominator",
			input1:      "-222.111",
			input2:      "-111.222",
			result:      lessThan,
		},
		{
			description: "Same sign, smaller numerator, same denominator",
			input1:      "111.222",
			input2:      "222.222",
			result:      lessThan,
		},
		{
			description: "Same sign, smaller numerator, same denominator",
			input1:      "-111.222",
			input2:      "-222.222",
			result:      greaterThan,
		},
		{
			description: "Same sign, smaller numerator, larger denominator",
			input1:      "111.333",
			input2:      "222.222",
			result:      lessThan,
		},
		{
			description: "Same sign, smaller numerator, larger denominator",
			input1:      "-111.333",
			input2:      "-222.222",
			result:      greaterThan,
		},
		{
			description: "Same sign, smaller numerator, smaller denominator",
			input1:      "111.111",
			input2:      "222.222",
			result:      lessThan,
		},
		{
			description: "Same sign, smaller numerator, smaller denominator",
			input1:      "-111.111",
			input2:      "-222.222",
			result:      greaterThan,
		},
		{
			description: "Different sign, same numerator, same denominator",
			input1:      "111.222",
			input2:      "-111.222",
			result:      greaterThan,
		},
		{
			description: "Different sign, same numerator, same denominator",
			input1:      "-111.222",
			input2:      "111.222",
			result:      lessThan,
		},
		{
			description: "Different sign, same numerator, larger denominator",
			input1:      "111.333",
			input2:      "-111.222",
			result:      greaterThan,
		},
		{
			description: "Different sign, same numerator, larger denominator",
			input1:      "-111.333",
			input2:      "111.222",
			result:      lessThan,
		},
		{
			description: "Different sign, same numerator, smaller denominator",
			input1:      "111.111",
			input2:      "-111.222",
			result:      greaterThan,
		},
		{
			description: "Different sign, same numerator, smaller denominator",
			input1:      "-111.111",
			input2:      "111.222",
			result:      lessThan,
		},
		{
			description: "Different sign, larger numerator, same denominator",
			input1:      "222.222",
			input2:      "-111.222",
			result:      greaterThan,
		},
		{
			description: "Different sign, larger numerator, same denominator",
			input1:      "-222.222",
			input2:      "111.222",
			result:      lessThan,
		},
		{
			description: "Different sign, larger numerator, larger denominator",
			input1:      "222.333",
			input2:      "-111.222",
			result:      greaterThan,
		},
		{
			description: "Different sign, larger numerator, larger denominator",
			input1:      "-222.333",
			input2:      "111.222",
			result:      lessThan,
		},
		{
			description: "Different sign, larger numerator, smaller denominator",
			input1:      "222.111",
			input2:      "-111.222",
			result:      greaterThan,
		},
		{
			description: "Different sign, larger numerator, smaller denominator",
			input1:      "-222.111",
			input2:      "111.222",
			result:      lessThan,
		},
		{
			description: "Different sign, smaller numerator, same denominator",
			input1:      "111.222",
			input2:      "-222.222",
			result:      greaterThan,
		},
		{
			description: "Different sign, smaller numerator, same denominator",
			input1:      "-111.222",
			input2:      "222.222",
			result:      lessThan,
		},
		{
			description: "Different sign, smaller numerator, larger denominator",
			input1:      "111.333",
			input2:      "-222.222",
			result:      greaterThan,
		},
		{
			description: "Different sign, smaller numerator, larger denominator",
			input1:      "-111.333",
			input2:      "222.222",
			result:      lessThan,
		},
		{
			description: "Different sign, smaller numerator, smaller denominator",
			input1:      "111.111",
			input2:      "-222.222",
			result:      greaterThan,
		},
		{
			description: "Different sign, smaller numerator, smaller denominator",
			input1:      "-111.111",
			input2:      "222.222",
			result:      lessThan,
		},
	}

	for _, test := range tests {
		d1, err := ParseDecimal(test.input1)
		if err != nil {
			t.Errorf("%s (input '%s'): expected success, received error '%v'.", test.description, test.input1, err)
			continue
		}
		d2, err := ParseDecimal(test.input2)
		if err != nil {
			t.Errorf("%s (input '%s'): expected success, received error '%v'.", test.description, test.input2, err)
			continue
		}

		result := d1.Cmp(d2)
		if test.result != result {
			t.Errorf("%s (comparing '%s' to '%s'): expected %d (%s), received %d (%s).", test.description, test.input1, test.input2, test.result, resultString[test.result], result, resultString[result])
		}
	}
}

type testResult struct {
	shouldFail, negative bool
	output               string
}
type operationTest struct {
	description, input1, input2 string
	result                      testResult
}

func testOperation(t *testing.T, tests []operationTest, op string) {
	var debugOp string
	switch op {
	case "+":
		debugOp = "adding"
	case "-":
		debugOp = "subtracting"
	default:
		t.Fatalf("Unsupported operation '%s'.", op)
	}

	for _, test := range tests {
		d1, err := ParseDecimal(test.input1)
		if err != nil {
			t.Errorf("%s (input '%s'): expected success, received error '%v'.", test.description, test.input1, err)
			continue
		}
		d2, err := ParseDecimal(test.input2)
		if err != nil {
			t.Errorf("%s (input '%s'): expected success, received error '%v'.", test.description, test.input2, err)
			continue
		}

		switch op {
		case "+":
			err = d1.Add(d2)
		case "-":
			err = d1.Sub(d2)
		}
		if err != nil {
			if !test.result.shouldFail {
				t.Errorf("%s (%s '%s' and '%s'): expected success, received error '%v'.", test.description, debugOp, test.input1, test.input2, err)
			}
			continue
		}
		if test.result.shouldFail {
			t.Errorf("%s (%s '%s' and '%s'): expected failure.", test.description, debugOp, test.input1, test.input2)
			continue
		}
		if test.result.negative && !d1.Negative {
			t.Errorf("%s (%s '%s' and '%s'): expected negative value.", test.description, debugOp, test.input1, test.input2)
		} else if !test.result.negative && d1.Negative {
			t.Errorf("%s (%s '%s' and '%s'): expected positive value.", test.description, debugOp, test.input1, test.input2)
		}
		if test.result.output != d1.String() {
			t.Errorf("%s (%s '%s' and '%s'): expected '%s', received '%s'.", test.description, debugOp, test.input1, test.input2, test.result.output, d1.String())
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []operationTest{
		{
			description: "Positive plus negative, result is zero",
			input1:      "111.111",
			input2:      "-111.111",
			result: testResult{
				output: "0.0",
			},
		},
		{
			description: "Negative plus positive, result is zero",
			input1:      "-111.111",
			input2:      "111.111",
			result: testResult{
				output: "0.0",
			},
		},
		{
			description: "Positive plus positive, sign stays the same",
			input1:      "111.111",
			input2:      "222.222",
			result: testResult{
				output: "333.333",
			},
		},
		{
			description: "Negative plus negative, sign stays the same",
			input1:      "-111.111",
			input2:      "-222.222",
			result: testResult{
				negative: true,
				output:   "-333.333",
			},
		},
		{
			description: "Positive plus negative, sign stays the same",
			input1:      "222.222",
			input2:      "-111.111",
			result: testResult{
				output: "111.111",
			},
		},
		{
			description: "Negative plus positive, sign stays the same",
			input1:      "-222.222",
			input2:      "111.111",
			result: testResult{
				negative: true,
				output:   "-111.111",
			},
		},
		{
			description: "Positive plus negative, sign becomes negative",
			input1:      "111.111",
			input2:      "-222.222",
			result: testResult{
				negative: true,
				output:   "-111.111",
			},
		},
		{
			description: "Negative plus positive, sign becomes positive",
			input1:      "-111.111",
			input2:      "222.222",
			result: testResult{
				output: "111.111",
			},
		},
		{
			description: "Positive plus positive, denominator carries into numerator",
			input1:      "111.555",
			input2:      "111.666",
			result: testResult{
				output: "223.221",
			},
		},
		{
			description: "Negative plus negative, denominator carries into numerator",
			input1:      "-111.555",
			input2:      "-111.666",
			result: testResult{
				negative: true,
				output:   "-223.221",
			},
		},
		{
			description: "Positive plus positive, uneven length denominators",
			input1:      "111.555",
			input2:      "111.444999",
			result: testResult{
				output: "222.999999",
			},
		},
		{
			description: "Positive plus positive, uneven length denominators",
			input1:      "111.444999",
			input2:      "111.555",
			result: testResult{
				output: "222.999999",
			},
		},
		{
			description: "Negative plus negative, uneven length denominators",
			input1:      "-111.555",
			input2:      "-111.444999",
			result: testResult{
				negative: true,
				output:   "-222.999999",
			},
		},
		{
			description: "Negative plus negative, uneven length denominators",
			input1:      "-111.444999",
			input2:      "-111.555",
			result: testResult{
				negative: true,
				output:   "-222.999999",
			},
		},
		{
			description: "Positive plus negative, uneven length denominators",
			input1:      "111.555",
			input2:      "-111.444999",
			result: testResult{
				output: "0.110001",
			},
		},
		{
			description: "Negative plus positive, uneven length denominators",
			input1:      "-111.555",
			input2:      "111.444999",
			result: testResult{
				negative: true,
				output:   "-0.110001",
			},
		},
		{
			description: "Positive plus positive, denominators result in a carry",
			input1:      "111.66",
			input2:      "111.66",
			result: testResult{
				output: "223.32",
			},
		},
		{
			description: "Negative plus negative, denominators result in a carry",
			input1:      "-111.66",
			input2:      "-111.66",
			result: testResult{
				negative: true,
				output:   "-223.32",
			},
		},
		{
			description: "Positive plus positive, denominators carry and simplify",
			input1:      "111.555",
			input2:      "111.645",
			result: testResult{
				output: "223.2",
			},
		},
		{
			description: "Negative plus negative, denominators carry and simplify",
			input1:      "-111.555",
			input2:      "-111.645",
			result: testResult{
				negative: true,
				output:   "-223.2",
			},
		},
		{
			description: "Positive plus positive, denominators simplify",
			input1:      "111.555",
			input2:      "111.445",
			result: testResult{
				output: "223.0",
			},
		},
		{
			description: "Negative plus negative, denominators simplify",
			input1:      "-111.555",
			input2:      "-111.445",
			result: testResult{
				negative: true,
				output:   "-223.0",
			},
		},
		{
			description: "Erroneous carry that would lead to a divide by zero panic",
			input1:      "0.17446744073709551615",
			input2:      "0.01",
			result: testResult{
				output: "0.18446744073709551615",
			},
		},
		{
			description: "Erroneous carry that would lead to a divide by zero panic",
			input1:      "0.18446744073709551615",
			input2:      "0.0",
			result: testResult{
				output: "0.18446744073709551615",
			},
		},
		{
			description: "Bounds checking the numerator",
			input1:      "18446744073709551615.0",
			input2:      "1.0",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Bounds checking the numerator via carry",
			input1:      "18446744073709551615.5",
			input2:      "0.5",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Bounds checking the denominator",
			input1:      "0.18446744073709551615",
			input2:      "0.00000000000000000001",
			result: testResult{
				shouldFail: true,
			},
		},
	}

	testOperation(t, tests, "+")
}

func TestSub(t *testing.T) {
	tests := []operationTest{
		{
			description: "Positive minus positive, result is zero",
			input1:      "222.222",
			input2:      "222.222",
			result: testResult{
				output: "0.0",
			},
		},
		{
			description: "Positive minus positive, signs stay the same",
			input1:      "222.222",
			input2:      "111.111",
			result: testResult{
				output: "111.111",
			},
		},
		{
			description: "Negative minus negative, signs stay the same",
			input1:      "-222.222",
			input2:      "-111.111",
			result: testResult{
				negative: true,
				output:   "-333.333",
			},
		},
		{
			description: "Positive minus negative, signs stay the same",
			input1:      "222.222",
			input2:      "-111.111",
			result: testResult{
				output: "333.333",
			},
		},
		{
			description: "Negative minus positive, signs stay the same",
			input1:      "-222.222",
			input2:      "111.111",
			result: testResult{
				negative: true,
				output:   "-333.333",
			},
		},
		{
			description: "Positive minus positive, sign becomes negative",
			input1:      "111.111",
			input2:      "222.222",
			result: testResult{
				negative: true,
				output:   "-111.111",
			},
		},
		{
			description: "Positive minus positive, uneven length denominators",
			input1:      "222.222",
			input2:      "111.111999",
			result: testResult{
				output: "111.110001",
			},
		},
		{
			description: "Positive minus positive, uneven length denominators",
			input1:      "111.111999",
			input2:      "222.222",
			result: testResult{
				negative: true,
				output:   "-111.110001",
			},
		},
		{
			description: "Positive minus positive, denominator borrows from numerator",
			input1:      "111.111",
			input2:      "0.999",
			result: testResult{
				output: "110.112",
			},
		},
		{
			description: "Bounds checking the numerator",
			input1:      "-18446744073709551615.0",
			input2:      "-1.0",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Bounds checking the numerator via carry",
			input1:      "-18446744073709551615.5",
			input2:      "-0.5",
			result: testResult{
				shouldFail: true,
			},
		},
		{
			description: "Bounds checking the denominator",
			input1:      "-0.18446744073709551615",
			input2:      "-0.00000000000000000001",
			result: testResult{
				shouldFail: true,
			},
		},
	}

	testOperation(t, tests, "-")
}

func TestFormattedString(t *testing.T) {
	tests := map[string]string{
		"1.01":                                      "1.01",
		"12.01":                                     "12.01",
		"123.01":                                    "123.01",
		"1234.01":                                   "1,234.01",
		"12345.01":                                  "12,345.01",
		"123456.01":                                 "123,456.01",
		"1234567.01":                                "1,234,567.01",
		"12345678.01":                               "12,345,678.01",
		"123456789.01":                              "123,456,789.01",
		"1234567890.01":                             "1,234,567,890.01",
		"18446744073709551615.18446744073709551615": "18,446,744,073,709,551,615.18446744073709551615",
		"-1.01":                                      "-1.01",
		"-12.01":                                     "-12.01",
		"-123.01":                                    "-123.01",
		"-1234.01":                                   "-1,234.01",
		"-12345.01":                                  "-12,345.01",
		"-123456.01":                                 "-123,456.01",
		"-1234567.01":                                "-1,234,567.01",
		"-12345678.01":                               "-12,345,678.01",
		"-123456789.01":                              "-123,456,789.01",
		"-1234567890.01":                             "-1,234,567,890.01",
		"-18446744073709551615.18446744073709551615": "-18,446,744,073,709,551,615.18446744073709551615",
	}

	for input, output := range tests {
		decimal, err := ParseDecimal(input)
		if err != nil {
			t.Errorf("'%s': Expected success, received error '%v'.", input, err)
			continue
		}
		if decimal.FormattedString() != output {
			t.Errorf("'%s': Expected '%s', received '%s'.", input, output, decimal.FormattedString())
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs()
	d1, _ := ParseDecimal("123456789.012345")
	d2, _ := ParseDecimal("8675309.1337")
	for i := 0; i < b.N; i++ {
		if err := d1.Add(d2); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSub(b *testing.B) {
	b.ReportAllocs()
	d1, _ := ParseDecimal("123456789.012345")
	d2, _ := ParseDecimal("8675309.1337")
	for i := 0; i < b.N; i++ {
		if err := d1.Sub(d2); err != nil {
			b.Fatal(err)
		}
	}
}
