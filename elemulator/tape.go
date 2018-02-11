package elemulator

import (
	"fmt"
	"math/rand"
)

// Tape is a binary tape that can extend infinitely
type Tape struct {
	data []uint64
}

// NewTape returns a new blank tape
func NewTape(src ...uint64) *Tape {
	return &Tape{src}
}

// NewRandTape return a new random tape
func NewRandTape(chunks int) *Tape {
	r := rand.New(rand.NewSource(99))
	data := make([]uint64, chunks)
	for i := 0; i < chunks; i++ {
		data[i] = r.Uint64()
	}
	return &Tape{data}
}

// Compare returns true if the contents of the tape are equivalent
func Compare(left *Tape, right *Tape) bool {
	if len(left.data) > len(right.data) { // swap so that left is the shorter
		left, right = right, left
	}

	for i := 0; i < len(left.data); i++ {
		if left.data[i] != right.data[i] {
			return false
		}
	}

	for i := len(left.data); i < len(right.data); i++ {
		if right.data[i] != 0 {
			return false
		}
	}

	return true
}

// String returns the string rep of a tape
func (t *Tape) String() string {
	s := ""
	for i := len(t.data) - 1; i >= 0; i-- {
		s += fmt.Sprintf("%064b", t.data[i])
	}
	return s
}
