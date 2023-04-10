package puzzles_part1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func compareToml() {
	result := []string{}

	file, err := os.Open("./disableInput")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	referenceFile, referenceEvents := getFileNameAndEvents(result[0])
	for _, line := range result[1:] {
		fileName, events := getFileNameAndEvents(line)
		missingReference, missingTarget := diffStrings(referenceEvents, events)
		fmt.Printf("Comparing %s vs %s\n%s has %v\n%s has %v\n-----\n", referenceFile, fileName, referenceFile, missingTarget, fileName, missingReference)
	}
}

func getFileNameAndEvents(line string) (string, []string) {
	result := []string{}
	for _, r := range strings.Split(line[strings.Index(line, "["):strings.Index(line, "]")], ",") {
		result = append(result, strings.TrimSpace(r))
	}
	return line[0:strings.Index(line, ":")], result
}

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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
