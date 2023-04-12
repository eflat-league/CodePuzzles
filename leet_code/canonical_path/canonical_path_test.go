package canonical_path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name  string
		input string
		exp   string
	}{
		{
			name:  "/",
			input: "/",
			exp:   "/",
		},
		{
			name:  "////",
			input: "////",
			exp:   "/",
		},
		{
			name:  "/hello/",
			input: "/hello/",
			exp:   "/hello",
		},
		{
			name:  "/hello/../",
			input: "/hello/../",
			exp:   "/",
		},
		{
			name:  "/../",
			input: "/../",
			exp:   "/",
		},
		{
			name:  "/home//foo/",
			input: "/home//foo/",
			exp:   "/home/foo",
		},
		{
			name:  "/home//foo/../../",
			input: "/home//foo/../../",
			exp:   "/",
		},
		{
			name:  "/./home//foo/../.././hello/.././durp/",
			input: "/./home//foo/../.././hello/.././durp/",
			exp:   "/durp",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := simplifyPath(test.input)
			assert.Equal(t, test.exp, got)
		})
	}
}
