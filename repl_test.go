package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "these are words",
			expected: []string{"these", "are", "words"},
		},
		{
			input:    "test string",
			expected: []string{"test", "string"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("\nExpected: %v\nActual: %v", expectedWord, word)
			}
		}
	}
}
