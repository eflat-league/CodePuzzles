/* Regex matching in this code is greedy.
 */
package operators

import "strings"

type RepeatTimes int

const (
	ONCE RepeatTimes = iota
	ONE_OR_MORE
	ZERO_OR_MORE
)

type RegexOperand struct {
	MatchValue string
	Repeat     RepeatTimes
}

// If it finds a match (greedy), returns true and remaining unmatched string.
// If it does not find a match, returns false and unmodified input.
func (r *RegexOperand) Evaluate(input string) (bool, string) {
	switch r.Repeat {
	case ONCE:
		return hardMatch(input, r.MatchValue)
	}
	return false, input
}

func hardMatch(input string, matchValue string) (bool, string) {
	if matchValue == "" {
		return true, input
	}

	result := input
	ok := true
	match := false
	for ok {
		result, ok = strings.CutPrefix(input, matchValue)
		match = true
	}

	return match, result
}
