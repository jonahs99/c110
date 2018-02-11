// Package elemulator contains an emulator for elementary cellular automata
package elemulator

// Iterate returns a single iteration of rule110
func Iterate(src *Tape) *Tape {
	in := src.data
	grow := in[len(in)-1]>>63 == 1

	x, y, z := make([]uint64, len(in)), in, make([]uint64, len(in))

	for i := 0; i < len(y)-1; i++ {
		x[i] = y[i]>>1 + (y[i+1]&1)<<63
	}
	x[len(y)-1] = y[len(y)-1] >> 1

	z[0] = y[0] << 1
	for i := 1; i < len(y); i++ {
		z[i] = y[i]<<1 + (y[i-1] >> 63 & 1)
	}

	outLen := len(in)
	if grow {
		outLen++
	}

	out := make([]uint64, outLen)
	for i := 0; i < len(y); i++ {
		out[i] = (^x[i] | ^y[i] | ^z[i]) & (y[i] | z[i])
	}

	if grow {
		out[len(y)] = 1
	}

	return &Tape{out}
}
