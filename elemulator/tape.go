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

// String returns the string rep of a tape
func (t Tape) String() string {
	s := ""
	for i := len(t.data) - 1; i >= 0; i-- {
		s += fmt.Sprintf("%064b", t.data[i])
	}
	return s
}
