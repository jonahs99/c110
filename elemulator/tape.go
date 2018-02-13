package elemulator

import (
	"fmt"
	"math/rand"
)

// Tape is a binary tape that can extend infinitely
type Tape struct {
	left, right int
	dataL       []uint64
	dataR       []uint64
}

// SetBounds sets bounds
func (tape *Tape) SetBounds(left, right int) {
	tape.left, tape.right = left, right
}

func (tape *Tape) leftChunks() int {
	if tape.left == 0 {
		return 0
	}
	return (tape.left-1)/64 + 1
}

func (tape *Tape) rightChunks() int {
	if tape.right == 0 {
		return 0
	}
	return (tape.right-1)/64 + 1
}

// NewTape returns a new blank tape
func NewTape(length int, src ...uint64) *Tape {
	return &Tape{
		0, length,
		make([]uint64, 0),
		src,
	}
}

// NewRandTape return a new random tape
func NewRandTape(chunks int) *Tape {
	r := rand.New(rand.NewSource(99))
	dataL := make([]uint64, chunks)
	dataR := make([]uint64, chunks)
	for i := 0; i < chunks; i++ {
		dataL[i] = r.Uint64()
		dataR[i] = r.Uint64()
	}
	return &Tape{
		chunks*64 - 1, chunks*64 - 1,
		dataL, dataR,
	}
}

// AtLeft returns a masked uint64 representing the bit at left pos, assuming the tape is cyclical
func (tape *Tape) AtLeft(pos int) uint64 {
	if tape == nil {
		return 0
	}

	period := tape.left + tape.right
	iL := pos % period

	var chunk uint64
	var i uint
	if iL < tape.left {
		chunk = tape.dataL[iL/64]
		i = uint(iL) % 64
	} else {
		iR := -iL - 1 + period
		chunk = tape.dataR[iR/64]
		i = 63 - uint(iR)%64
	}

	mask := chunk & (1 << i)
	if mask == 0 {
		return 0
	}
	return 1 << uint(pos%64)
}

// AtRight is like AtLeft but right
func (tape *Tape) AtRight(pos int) uint64 {
	if tape == nil {
		return 0
	}

	period := tape.left + tape.right
	iR := pos % period

	var chunk uint64
	var i uint
	if iR < tape.right {
		chunk = tape.dataR[iR/64]
		i = 63 - uint(iR)%64
	} else {
		iL := -iR - 1 + period
		chunk = tape.dataL[iL/64]
		i = uint(iL) % 64
	}

	mask := chunk & (1 << i)
	if mask == 0 {
		return 0
	}
	return 1 << uint(63-pos%64)
}

// Compare returns true if the contents of the tape are equivalent
func Compare(left *Tape, right *Tape) bool {
	return false
}

// String returns the string rep of a tape
func (tape *Tape) String() string {
	s := ""
	for i := len(tape.dataL) - 1; i >= 0; i-- {
		s += fmt.Sprintf("%064b", tape.dataL[i])
	}
	s += "|"
	for i := 0; i < len(tape.dataR); i++ {
		s += fmt.Sprintf("%064b", tape.dataR[i])
	}
	return s
}
