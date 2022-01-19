// Counting Duplicates -- https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1/train/go
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
	test_fancy("abcdeaB11", 3)
	test_fancy("indivisibility", 1)
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

// fancy solution from submissions on code-wars
func duplicate_count_fancy(s1 string) (c int) {
	m := make(map[rune]int)
	for _, r := range strings.ToLower(s1) {
		if m[r]++; m[r] == 2 {
			c++
		}
	}
	return
}

func test_fancy(s string, e int) {
	p("provided string:", s, "expected result:", e, "fancy result provided:", duplicate_count_fancy(s))
}
