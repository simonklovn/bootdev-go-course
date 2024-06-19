package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		recipient string
		text      string
		expected  string
	}
	tests := []testCase{
		{
			recipient: "Honey Bunny",
			text:      "I love you, Pumpkin.",
			expected: `
To: Honey Bunny
Message: I love you, Pumpkin.
`,
		},
		{
			recipient: "Pumpkin",
			text:      "I love you, Honey Bunny.",
			expected: `
To: Pumpkin
Message: I love you, Honey Bunny.
`,
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				recipient: "poor sap 1",
				text:      "And you will know My name is the Lord when I lay My vengeance upon thee.",
				expected: `
To: poor sap 1
Message: And you will know My name is the Lord when I lay My vengeance upon thee.
`,
			},
			{
				recipient: "Fabienne",
				text:      "Zed's dead, baby. Zed's dead.",
				expected: `
To: Fabienne
Message: Zed's dead, baby. Zed's dead.
`,
			},
		}...)
	}

	for _, test := range tests {
		m := Message{Recipient: test.recipient, Text: test.text}
		messageText := getMessageText(m)
		if strings.TrimSpace(messageText) != strings.TrimSpace(test.expected) {
			t.Errorf(`Test Failed:
input:
* %v
* %v
=>
expected:
%v
actual:
%v
===========================
`,
				m.Recipient,
				m.Text,
				test.expected,
				messageText,
			)
		} else {
			fmt.Printf(`Test Passed:
input:
* %v
* %v
=>
expected:
%v
actual:
%v
===========================
`,
				m.Recipient,
				m.Text,
				test.expected,
				messageText,
			)
		}
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
