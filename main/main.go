package main

import (
	"c110/elemulator"
)

func main() {
	breadth := 1000
	time := 100000

	img := elemulator.NewImage(breadth, time)
	tape := elemulator.NewRandTape(breadth / 2 / 64)

	elemulator.DrawLine(img, tape, 0)
	for i := 1; i < time; i++ {
		tape = elemulator.Iterate(tape, tape, tape)
		go elemulator.DrawLine(img, tape, i)
	}

	elemulator.SavePNG(img, "out.png")
}
