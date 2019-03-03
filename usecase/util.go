package usecase

import (
	"image"
	"image/png"
	"os"
)

func readImageFile(fn string) (img image.Image, err error) {
	w, err := os.Open(fn)
	defer w.Close()
	if err != nil {
		return
	}
	return png.Decode(w)
}

func writeImageFile(fn string, img image.Image) (err error) {
	w, err := os.Create(fn)
	defer w.Close()
	if err != nil {
		return
	}
	return png.Encode(w, img)
}
