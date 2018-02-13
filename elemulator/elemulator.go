// Package elemulator contains an emulator for elementary cellular automata
package elemulator

// Iterate returns a single iteration of rule110
func Iterate(src *Tape, bgLeft *Tape, bgRight *Tape) *Tape {
	//left, right := src.left+1, src.right+1
	left, right := src.left, src.right
	out := Tape{
		left, right,
		nil, nil,
	}
	lchunks, rchunks := out.leftChunks(), out.rightChunks()
	chunks := lchunks + rchunks

	x := make([]uint64, chunks)
	y := make([]uint64, chunks)
	z := make([]uint64, chunks)

	for i := 0; i < lchunks; i++ {
		if i == 0 && len(src.dataL) < lchunks {
			continue
		}
		y[i] = src.dataL[lchunks-1-i]
	}
	for i := lchunks; i < chunks; i++ {
		if i == chunks-1 && len(src.dataR) < rchunks {
			continue
		}
		y[i] = src.dataR[i-lchunks]
	}

	for i := 0; i < chunks; i++ {
		x[i] = y[i] >> 1
		z[i] = y[i] << 1
		if i > 0 {
			z[i-1] += (y[i] >> 63) & 1
			x[i] += (y[i-1] & 1) << 63
		}
	}
	if bgLeft != nil {
		if bgLeft.AtLeft(src.left) > 0 {
			x[0] |= 1 << (uint(src.left+63) % 64)
		} else {
			x[0] &= ^(1 << (uint(src.left+63) % 64))
		}
	}
	if bgRight != nil {
		if bgLeft.AtRight(src.right) > 0 {
			z[chunks-1] |= 1 << ((64 - (uint(src.right) % 64)) % 64)
		} else {
			z[chunks-1] &= ^(1 << ((64 - (uint(src.right) % 64)) % 64))
		}
	}

	dataL := make([]uint64, lchunks)
	dataR := make([]uint64, rchunks)

	for i := 0; i < lchunks; i++ {
		dataL[lchunks-1-i] = (^x[i] | ^y[i] | ^z[i]) & (y[i] | z[i])
	}
	for i := lchunks; i < chunks; i++ {
		dataR[i-lchunks] = (^x[i] | ^y[i] | ^z[i]) & (y[i] | z[i])
	}

	out.dataL = dataL
	out.dataR = dataR

	return &out

}
