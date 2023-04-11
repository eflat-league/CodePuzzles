package remove_stars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveStars(t *testing.T) {
	tests := []struct {
		name  string
		input string
		exp   string
	}{
		{
			name:  "empty",
			input: "",
			exp:   "",
		},
		{
			name:  "basic",
			input: "ab**c*",
			exp:   "",
		},
		{
			name:  "basic non empty result",
			input: "hello*ab**c*bye",
			exp:   "hellbye",
		},
		{
			name:  "basic non empty result left",
			input: "o*ab**c*bye",
			exp:   "bye",
		},
		{
			name:  "basic non empty result right",
			input: "hello*ab**c*",
			exp:   "hell",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := removeStars(test.input)
			assert.Equal(t, test.exp, got)
		})
	}
}
