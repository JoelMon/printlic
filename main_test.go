// test suite
package cmd

import (
	"testing"
)

func TestInfoCheck(t *testing.T) {
	for _, test := range inputCheckTestCases {
		got, err := inputCheck(test.fname, test.lname, test.mname, test.birth, test.sex)

		if test.errorExpected {
			// check if err is of type error
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("FAIL: %s - Expected an error.", test.description)
			}
		} else {
			if err != nil {
				t.Fatalf("FAIL: %s - An error was not expected.\n Error msg: %q.", test.description, err)
			}
		}

		if got != test.expected {
			t.Fatalf("FAIL: %s - Expected %q but got %q", test.description, test.expected, got)
		}
		t.Logf("Passed - %s", test.description)
	}
}
