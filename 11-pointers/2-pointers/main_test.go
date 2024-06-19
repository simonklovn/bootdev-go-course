package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		messageIn string
		expected  string
	}
	tests := []testCase{
		{
			"English, motherfubber, do you speak it?",
			"English, mother****er, do you speak it?",
		},
		{
			"Oh man I've seen some crazy ass shiz in my time...",
			"Oh man I've seen some crazy ass **** in my time...",
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				"Does he look like a witch?",
				"Does he look like a *****?",
			},
		}...)
	}

	for _, test := range tests {
		original := test.messageIn
		removeProfanity(&test.messageIn)
		if test.messageIn != test.expected {
			t.Errorf(`Test Failed:
input:
%v
=>
expected:
%v
actual:
%v
===========================
`,
				original,
				test.expected,
				test.messageIn,
			)
		} else {
			fmt.Printf(`Test Passed:
  input:
%v
=>
  expected:
%v
  actual:
%v
===========================
`,
				original,
				test.expected,
				test.messageIn,
			)
		}
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
