// test suite
package cmd

import (
	"testing"
)

// TestMakeLower is someting
func TestInfoCheck(t *testing.T) {
	for _, test := range inputCheckTestCases {
		got, err := inputCheck(test.input)

		if test.errorExpected {
			// check if err is of type error
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("FAIL: %s - Expected an error when the input is %q.", test.description, test.input)
			}
		} else {
			if err != nil {
				t.Fatalf("FAIL: %s - An error was not expected with input %q.\n Error msg: %q.", test.description, test.input, err)
			}
		}

		if got != test.expected {
			t.Fatalf("FAIL: %s - Expected %q but got %q with input %q", test.description, test.expected, got, test.input)
		}
		t.Logf("Passed - %s", test.description)
	}
}
