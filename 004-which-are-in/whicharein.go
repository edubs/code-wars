// which are in -- https://www.codewars.com/kata/550554fd08b86f84fe000a58/train/go
package main

import (
	"fmt"
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string { // couldn't use (r []string) shortcut because of empty slice vs nil slice issue
	r := []string{}
	for _, string1 := range array1 {
		for _, string2 := range array2 {
			if strings.Contains(string2, string1) && !contains(r, string1) {
				r = append(r, string1)
			}
		}
	}
	sort.Strings(r)
	return r
}

func contains(s1 []string, w string) bool {
	for _, i := range s1 {
		if i == w {
			return true
		}
	}
	return false
}

func main() {

	a1 := []string{"arp", "live", "strong"}
	a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	fmt.Println(InArray(a1, a2))
	b1 := []string{""}
	b2 := []string{"asdf"}
	fmt.Println(InArray(b1, b2))
	a1 = []string{"cod", "code", "wars", "ewar", "ar"}
	a2 = []string{}
	fmt.Println(InArray(a1, a2))
	a1 = []string{"duplicates", "duplicates"}
	a2 = []string{"duplicates", "duplicates"}
	fmt.Println(InArray(a1, a2))
}
