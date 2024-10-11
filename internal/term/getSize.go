package term

import (
	"os"
	"golang.org/x/term"
)

func Size() (w, h int, err error) {
	termNumber := int(os.Stdin.Fd())
	w, h, err = term.GetSize(termNumber)
	if err != nil {
		return w, h, err
	}
	return w, h, nil
}
