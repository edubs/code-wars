// Highest and Lowest -- https://www.codewars.com/kata/554b4ac871d6813a03000035
package highlow

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var p = fmt.Println

func HighAndLow(in string) string {
	// slice input string on space delimeter and convert to int
	ints := []int{}
	for _, c := range strings.Split(in, " ") {
		i, err := strconv.Atoi(c)

		if err != nil {
			p("error:", err)
			os.Exit(1)
		}

		ints = append(ints, i)
	}

	// set both swap vars to the value at first index
	low := ints[0]
	high := ints[0]

	for _, j := range ints {
		if j < low {
			low = j
		}
		if j > high {
			high = j
		}
	}
	return strconv.Itoa(high) + " " + strconv.Itoa(low)
}
