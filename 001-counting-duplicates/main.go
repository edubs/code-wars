package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {
	test("abcde", 0)
	test("abcdea", 1)
	test("abcdeaB11", 3)
	test("indivisibility", 1)
	test("11", 1)
	test("aA", 1)
}

func duplicate_count(s1 string) int {

	lowerString := strings.ToLower(s1) // get a lowercase verion of the provided string

	matches := make(map[string]int) // create a map to store character matches (no duplicates)

	for _, character := range lowerString {
		char := string(character)
		count := strings.Count(lowerString, char)
		// check that the number of appearences is gt 1 and that we haven't counted it already
		if matches[char] == 0 && count > 1 {
			matches[char] = count
		}
	}
	return len(matches)
}

func test(s string, e int) {
	p("provided string:", s, "expected result:", e, "result provided:", duplicate_count(s))
}
