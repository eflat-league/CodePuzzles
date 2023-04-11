package valid_brackets

import "testing"
import "github.com/stretchr/testify/assert"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		exp   bool
	}{
		{
			name:  "empty",
			input: "",
			exp:   true,
		},
		{
			name:  "single pair",
			input: "()",
			exp:   true,
		},
		{
			name:  "invalid",
			input: ")(",
			exp:   false,
		},
		{
			name:  "invalid2",
			input: "{]",
			exp:   false,
		},
		{
			name:  "multiple pairs",
			input: "({[]})",
			exp:   true,
		},
		{
			name:  "multiple pairs 2",
			input: "[((((({[]})))))]",
			exp:   true,
		},
		{
			name:  "unclosed 1",
			input: "[",
			exp:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isValid(test.input)
			assert.Equal(t, test.exp, got)
		})
	}
}
