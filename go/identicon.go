package main

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	ChunkSize = 3
	GridSize  = 5
	CellWidth = 100
)

type Identicon struct {
	input    string
	checksum [16]byte
	grid     [GridSize][GridSize]byte
	colour   color.RGBA
	img      *image.RGBA
}

func NewIdenticon(input string) *Identicon {
	checksum := md5.Sum([]byte(input))
	colour := color.RGBA{checksum[0], checksum[1], checksum[2], 0xff}

	width := GridSize * CellWidth
	height := width
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	return &Identicon{
		input:    input,
		checksum: checksum,
		colour:   colour,
		img:      img,
	}
}

// generateGrid creates a 5x5 grid from the MD5 checksum
func (ic *Identicon) generateGrid() {
	for i := 0; i < GridSize; i++ {
		lower := i * ChunkSize
		upper := lower + ChunkSize
		bytes := [ChunkSize]byte(ic.checksum[lower:upper])

		ic.grid[i] = mirrorBytes(bytes)
	}
}

// generateImage creates the internal representation of the final image
// and writes it out to file.
func (ic *Identicon) generateImage() {
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}

	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if ic.grid[i][j]%2 == 0 {
				paintCell(ic.img, ic.colour, j, i)
			} else {
				paintCell(ic.img, white, j, i)
			}
		}
	}
}

// write will write out the image to the give filename
func (ic *Identicon) write(name string) {
	f, _ := os.Create(name)
	png.Encode(f, ic.img)
}

// String representation of the identicon
//
// Example:
//
//	·███·
//	█···█
//	█···█
//	██·██
//	█·█·█
func (ic *Identicon) String() string {
	s := ""
	for _, line := range ic.grid {
		for _, cell := range line {
			if cell%2 == 0 {
				s += "█"
			} else {
				s += "·"
			}
		}
		s += "\n"
	}

	return s
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("You need to provide input text on the command line.")
		return
	}

	input := os.Args[1]

	identicon := NewIdenticon(input)
	identicon.generateGrid()
	identicon.generateImage()
	// Spit out an ASCII representation as well, just for fun
	fmt.Println(identicon)

	identicon.write("identicon.png")
}

// mirrorBytes creates a mirrored version of the given bytes
// eg. [A B C] => [A B C B A]
func mirrorBytes(s [ChunkSize]byte) [GridSize]byte {
	return [GridSize]byte{
		s[0], s[1], s[2], s[1], s[0],
	}
}

// Paint a specific solid cell in the image
func paintCell(img *image.RGBA, colour color.RGBA, gridX, gridY int) {
	minX := gridX * CellWidth
	maxX := minX + CellWidth
	minY := gridY * CellWidth
	maxY := minY + CellWidth

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			img.Set(x, y, colour)
		}
	}
}
