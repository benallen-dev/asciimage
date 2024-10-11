package main

import (
	"fmt"
	"log"
	"os"

	"github.com/benallen-dev/asciimage/internal/img"
	"github.com/benallen-dev/asciimage/internal/term"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: asciimage <file path>")
		return
	}

	source, err := img.Read(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	w, h, err := term.Size()
	if err != nil {
		log.Fatal(err)
	}

	newImg := img.Resize(source, w, h)
	img.Draw(newImg)
}
