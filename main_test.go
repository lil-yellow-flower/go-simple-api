package main

import (
	"reflect"
	"testing"
)

func TestSorter(t *testing.T) {
	testCase := []string{"HEL-GOT", "GOT-ARL", "CPH-HEL"}
	got := sorter(testCase)
	want := []string{"CPH-HEL", "HEL-GOT", "GOT-ARL"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}
}
