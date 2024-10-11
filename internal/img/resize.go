package img

import (
	"image"
	"math"

	"github.com/nfnt/resize"
)

func Resize(img image.Image, w, h int) image.Image {
	// The command prompt takes two lines, so we need to subtract 2 from the height
	h = h - 2

	s := 1.0
	w = int(math.Floor(float64(w)*s))
	h = int(math.Floor(float64(h)*s))

	// Set the largest of w,h to 0 to resize inside the smallest size
	if (h > w) {
		h = 0
	} else {
		w = 0
	}

	// TODO: Fix how the resize assumes the scaling by the smallest side
	// of the image means the other side fits inside the terminal.

	return resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
}
