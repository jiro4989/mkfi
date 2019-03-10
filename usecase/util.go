package usecase

import (
	"image"
	"image/png"
	"os"

	"github.com/jiro4989/mkfi/log"
)

func readImageFile(fn string) (img image.Image, err error) {
	log.Debug("read image file. image=", fn)
	w, err := os.Open(fn)
	defer w.Close()
	if err != nil {
		return
	}
	log.Debug("read success.")
	return png.Decode(w)
}

func writeImageFile(fn string, img image.Image) (err error) {
	log.Debug("write image file. image=", fn)
	w, err := os.Create(fn)
	defer w.Close()
	if err != nil {
		return
	}
	log.Debug("write success.")
	return png.Encode(w, img)
}
