package elemulator

import (
	"testing"
)

func TestIterate(t *testing.T) {
	// 1, 3, 7, 13, 31, 49, 115, 215, 509, 775, 1805
	cases := []struct {
		in, want uint64
	}{
		{0, 0},
		{1, 3},
		{3, 7},
		{7, 13},
		{13, 31},
		{31, 49},
		{49, 115},
		{115, 215},
		{215, 509},
		{509, 775},
	}
	for _, c := range cases {
		got := Iterate(NewTape(c.in), nil, nil).dataR[0]
		if got != c.want {
			t.Errorf("Iterate(%b) == %b, want %b", c.in, got, c.want)
		}
	}
}

func TestIterateGrow(t *testing.T) {
	in := NewTape(0, uint64(3)<<62)
	want := NewTape(0, uint64(3)<<62, 1)

	got := Iterate(in, nil, nil)
	if !Compare(got, want) {
		t.Errorf("Iterate(%v) == %v, want %v", in, got, want)
	}
}
