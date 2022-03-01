package main

import (
	"reflect"
	"testing"
)

func TestPrimeGap(t *testing.T) {
	tables := []struct {
		g    int
		m    int
		n    int
		want []int
	}{
		{2, 5, 7, []int{5, 7}},
		{2, 100, 110, []int{101, 103}},
		{4, 100, 110, []int{103, 107}},
		{6, 100, 110, nil},
		{8, 300, 400, []int{359, 367}},
		{10, 300, 400, []int{337, 347}},
	}

	for _, table := range tables {
		got := Gap(table.g, table.m, table.n)
		if !reflect.DeepEqual(got, table.want) {
			t.Errorf("Wanted: %v Got: %v", table.want, got)
		}
	}
}
