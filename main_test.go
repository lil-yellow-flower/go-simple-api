package main

import (
	"reflect"
	"testing"
)

func TestSorterSmall(t *testing.T) {
	testCase := []string{"HEL-GOT", "GOT-ARL", "CPH-HEL"}
	got := sorter(testCase)
	want := []string{"CPH-HEL", "HEL-GOT", "GOT-ARL"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}
}

func TestSorterBig(t *testing.T) {
	testCase := []string{
		"TKY-HGK", "HGK-SYD", "THR-ISR", "ARL-OSL", "GMB-KNY", "HEL-GOT", "SAF-SAM", "NYZ-TAI",
		"SAM-BRS", "ARG-CHI", "VEN-PAN", "OSL-NYC", "CLF-TKY", "TAI-VIE", "GOT-ARL", "KZK-AFG",
		"CPH-HEL", "VIE-MSW", "NYC-CLF", "CUB-MIA", "BNG-THR", "CHI-VEN", "MSW-KZK", "AFG-PKI",
		"SYD-NYZ", "PKI-BNG", "PAN-CUB", "KNY-SAF", "BRS-ARG", "ISR-GMB", "MIA-SWE"}

	got := sorter(testCase)
	want := []string{
		"CPH-HEL", "HEL-GOT", "GOT-ARL", "ARL-OSL", "OSL-NYC", "NYC-CLF", "CLF-TKY", "TKY-HGK",
		"HGK-SYD", "SYD-NYZ", "NYZ-TAI", "TAI-VIE", "VIE-MSW", "MSW-KZK", "KZK-AFG", "AFG-PKI",
		"PKI-BNG", "BNG-THR", "THR-ISR", "ISR-GMB", "GMB-KNY", "KNY-SAF", "SAF-SAM", "SAM-BRS",
		"BRS-ARG", "ARG-CHI", "CHI-VEN", "VEN-PAN", "PAN-CUB", "CUB-MIA", "MIA-SWE"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}
}
