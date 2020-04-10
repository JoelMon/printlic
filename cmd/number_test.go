// test suite
package cmd

import (
	"testing"
)

// TestMakeLower is someting
func TestInfoCheck(t *testing.T) {
	for _, test := range inputCheckTestCases {
		err := inputCheck(test.firstName, test.lastName, test.middleName, test.birth, test.sex)

		if test.errorExpected {
			// check if err is of type error
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("FAIL: %s - Expected an error", test.description)
			}
		} else {
			if err != nil {
				t.Fatalf("FAIL: %s - An error was not expected.\n Error msg: %q.", test.description, err)
			}
		}

		t.Logf("Passed - %q with error message: %v", test.description, err)
	}
}

// TestEFirstName tests the eFirstName function
func TestEFirstName(t *testing.T) {
	for _, test := range eFirstNameTestCases {
		got, err := eFirstName(test.firstName, test.middleName)

		if test.errorExpected {
			// check if err is of type error
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("FAIL: %s - Expected an error", test.description)
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

// TestPad tests the pad function
func TestPad(t *testing.T) {
	for _, test := range padTestCases {
		got := pad(test.codedName)

		if got != test.expected {
			t.Fatalf("FAIL: %s - Expected %q but got %q", test.description, test.expected, got)
		}
		t.Logf("Passed - %s", test.description)
	}
}
