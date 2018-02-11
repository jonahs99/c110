package main

import (
	"c110/elemulator"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(99))

	img := elemulator.NewImage(4000, 2000)
	tape := elemulator.NewTape(r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64())

	for i := 0; i < 10000; i++ {
		tape = elemulator.Iterate(tape)
		elemulator.DrawLine(img, tape, i)
	}

	elemulator.SavePNG(img, "out.png")
}
