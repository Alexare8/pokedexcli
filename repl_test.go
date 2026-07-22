package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Incorrect number of words. Actual: %v, Expected: %v", len(actual), len(c.expected))
			continue
		}
		for i, word := range actual {
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("'%v' does not match expected '%v'", word, expectedWord)
			}
		}
	}
}
