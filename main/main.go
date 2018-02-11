package main

import (
	"c110/elemulator"
)

func main() {
	img := elemulator.NewImage(400, 400)
	tape := elemulator.NewRandTape(400/64 + 1)

	for i := 0; i < 400; i++ {
		tape = elemulator.Iterate(tape)
		elemulator.DrawLine(img, tape, i)
	}

	elemulator.SavePNG(img, "out.png")
}
