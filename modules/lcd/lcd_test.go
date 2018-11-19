package lcd

import (
	"testing"
)

func TestLCD(t *testing.T) {
	tcs := []struct {
		in  int
		out string
	}{
		{0, " _ \n| |\n|_|"},
		{1, "   \n  |\n  |"},
		{2, " _ \n _|\n|_ "},
		{3, " _ \n _|\n _|"},
		{4, "   \n|_|\n  |"},
		{5, " _ \n|_ \n _|"},
		{6, " _ \n|_ \n|_|"},
		{7, " _ \n  |\n  |"},
		{8, " _ \n|_|\n|_|"},
		{9, " _ \n|_|\n  |"},
		{123, "     _   _ \n  |  _|  _|\n  | |_   _|"},
	}
	for _, tt := range tcs {
		r := LCD(tt.in)
		t.Logf("\n%s\n", r)
		if r != tt.out {
			t.Errorf("expect \n%s\n got \n%s\n", tt.out, r)
		}
	}
}
