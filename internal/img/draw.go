package img

import (
	"fmt"
	"image"
)

func Draw(img image.Image) {
	s := img.Bounds().Size()
	w := s.X
	h := s.Y

	for y := 0; y < h; y = y + 2 {
		for x := 0; x < w; x++ {
			rt, gt, bt, _ := img.At(x, y).RGBA()
			rt = rt >> 8 // Move color to 8bpp
			gt = gt >> 8
			bt = bt >> 8

			rb, gb, bb, _ := img.At(x, y+1).RGBA()
			rb = rb >> 8 // Move color to 8bpp
			gb = gb >> 8
			bb = bb >> 8

			fmt.Printf("\033[48;2;%d;%d;%dm\033[38;2;%d;%d;%dm▄\033[0m", rt, gt, bt, rb, gb, bb)
		}
		fmt.Println()
	}
}
