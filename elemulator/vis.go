package elemulator

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// NewImage returns a new image.Gray with the specified size
func NewImage(w, h int) *image.Gray {
	return image.NewGray(image.Rect(0, 0, w, h))
}

// SavePNG saves the image.Gray to the path
func SavePNG(img *image.Gray, path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Printf("Error creating the image file.\n")
		return
	}
	png.Encode(file, img)
}

// DrawLine draws the tape at the specified index on the image
func DrawLine(img *image.Gray, line *Tape, index int) {
	bounds := img.Bounds()
	w := bounds.Max.X
	center := w / 2

Left:
	for i := 0; i < line.leftChunks(); i++ {
		chunk := line.dataL[i]
		for x := 63; x >= 0; x-- {
			pix := center - 1 - 64*i - x
			if pix < 0 {
				break Left
			}
			if (chunk>>uint(x))&1 == 1 {
				img.SetGray(pix, index, color.Gray{0})
			} else {
				img.SetGray(pix, index, color.Gray{255})
			}
		}
	}
Right:
	for i := 0; i < line.rightChunks(); i++ {
		chunk := line.dataR[i]
		for x := 0; x < 64; x++ {
			pix := center + 64*(i+1) - 1 - x
			if pix >= w {
				break Right
			}
			if (chunk>>uint(x))&1 == 1 {
				img.SetGray(pix, index, color.Gray{0})
			} else {
				img.SetGray(pix, index, color.Gray{255})
			}
		}
	}

}
