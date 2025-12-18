package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {

	tests := map[string]struct {
		input    string
		expected []string
	}{
		"two words plus surounding spaces": {input: " hello world ",
			expected: []string{"hello", "world"}},
		"one word": {input: " helloworld ",
			expected: []string{"helloworld"}},
		"empty string": {input: "",
			expected: []string{""}},
		"should lowecase everything": {input: "HelLO wOrLD",
			expected: []string{"hello world"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("expected: %v got: %v", tc.expected, got)
			}
		})

	}

}
