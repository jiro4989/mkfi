package infrastructure

import (
	"image"

	"github.com/julienschmidt/httprouter"
	"github.com/oliamb/cutter"
)

type (
	HTTPParams httprouter.Params
	CropConfig cutter.Config
)

const ImageTopLeft = cutter.TopLeft

func CropImage(src image.Image, c CropConfig) (image.Image, error) {
	return cutter.Crop(src, c)
}
