package img

import (
	"fmt"
	"image"
)

func Draw(img image.Image) {
	s := img.Bounds().Size()
	w := s.X
	h := s.Y

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r = r >> 8 // Move color to 8bpp
			g = g >> 8
			b = b >> 8

			fmt.Printf("\033[38;2;%d;%d;%dm██\033[0m", r, g, b)
		}
		fmt.Println()
	}
}
