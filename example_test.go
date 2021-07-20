package gomodexample

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty string", input: "", want: ""},
		{name: "one symbol", input: "d", want: "d"},
		{name: "even string", input: "string", want: "gnirts"},
		{name: "odd string", input: "hello", want: "olleh"},
		{name: "russion odd string", input: "нет", want: "тен"},
		{name: "russion even string", input: "привет", want: "тевирп"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := ReverseString(test.input)
			if test.want != got {
				t.Fatalf("expected: %s, got: %s", test.input, got)
			}
		})
	}
}
