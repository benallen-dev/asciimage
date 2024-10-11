package img

import (
	"bytes"
	"image"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	_ "image/png"
	_ "image/jpeg"
)

func Read(filename string) (img image.Image, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)
	img, _, err = image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return img, nil
}
