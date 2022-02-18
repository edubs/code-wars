package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrimeGap(t *testing.T) {
	got := Gap(2, 3, 50)
	want := []int{3, 5}

	if Converter(got) != Converter(want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func Converter(i []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(i)), ", "), "[]")
}
