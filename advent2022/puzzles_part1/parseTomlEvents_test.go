package puzzles_part1

import "testing"

func TestCompareToml(t *testing.T) {
	compareToml()
}

//run this to see difference between

// returns a list of things in b but not in a, and a list of things in a but not in b, in that order
func diffStrings(a, b []string) ([]string, []string) {
	missingFromA := []string{}
	missingFromB := []string{}
	for _, aEl := range a {
		if !contains(b, aEl) {
			missingFromB = append(missingFromB, aEl)
		}
	}
	for _, bEl := range b {
		if !contains(a, bEl) {
			missingFromA = append(missingFromA, bEl)
		}
	}
	return missingFromA, missingFromB
}
