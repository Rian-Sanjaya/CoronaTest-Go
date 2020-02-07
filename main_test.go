package main

import "testing"

func TestMatchDNA(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"abbab ba", "1 2"},
		{"hello world", "No Match!"},
		{"banana nan", "0 2"},
		{"cgatcg gc", "1 3"},
		{"atcgatcga cgg", "2 6"},
		{"aardvark ab", "0 1 5"},
	}

	for _, test := range tests {
		if output := matchDNA(test.input); output != test.expected {
			t.Error("Test Failed: inputted: " + test.input + ", expected: " + test.expected + ", recieved: " + output)
		}
	}
}
