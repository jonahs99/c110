package elemulator

import (
	"strings"
	"testing"
)

func TestNewTape(t *testing.T) {
	NewTape(0)
	NewTape(0, 1)
}

func TestNewRandTape(t *testing.T) {
	testLen := 10
	tape := NewRandTape(testLen)

	got := len(tape.dataL)
	if got != testLen {
		t.Errorf("len(NewRandTape(%d).dataL) == %d, want %d", testLen, got, testLen)
	}
}

func TestCompare(t *testing.T) {
	cases := []struct {
		left, right *Tape
		want        bool
	}{
		{NewTape(0), NewTape(0), true},
		{NewTape(0), NewTape(1), false},
		{NewTape(137), NewTape(137), true},
		{NewTape(137, 0, 0), NewTape(137), true},
		{NewTape(137), NewTape(137, 0, 0), true},
		{NewTape(137, 0, 1), NewTape(137), false},
		{NewTape(0, 137), NewTape(137), false},
	}

	for _, c := range cases {
		left, right, want := c.left, c.right, c.want
		got := Compare(left, right)
		if got != want {
			t.Errorf("Compare(%v, %v) == %t, want %t", left, right, got, want)
		}
	}
}

func TestString(t *testing.T) {
	in := NewTape(0, 0xf)
	want := strings.Repeat("0", 60) + "1111" + strings.Repeat("0", 64)

	got := in.String()
	if got != want {
		t.Errorf("tape.String() == %s, want %s", got, want)
	}
}
