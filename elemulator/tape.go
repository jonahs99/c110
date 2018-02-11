package elemulator

import (
	"fmt"
)

// Tape is a binary tape that can extend infinitely
type Tape struct {
	data []uint64
}

// NewTape returns a new tape
func NewTape(src ...uint64) Tape {
	return Tape{src}
}

// String returns the string rep of a tape
func (t Tape) String() string {
	s := ""
	for i := len(t.data) - 1; i >= 0; i-- {
		s += fmt.Sprintf("%064b", t.data[i])
	}
	return s
}
