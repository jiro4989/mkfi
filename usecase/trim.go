package usecase

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"path/filepath"
	"sync"

	"github.com/jiro4989/mkfi/domain"
	"github.com/oliamb/cutter"
)

func TrimImageFiles(rect domain.Rectangle, outDir string, targetFiles []string) ([]string, error) {
	return logic(rect, outDir, trimImageFile, targetFiles)
}

func trimImageFile(wg *sync.WaitGroup, q chan string, in1 interface{}, outDir string, createdFiles []string, errs []error) {
	defer wg.Done()

	rect, ok := in1.(domain.Rectangle)
	if !ok {
		err := errors.New(fmt.Sprintf("illegal input data. input=%v", in1))
		errs = append(errs, err)
		return
	}
	var (
		x   = rect.X
		y   = rect.Y
		w   = rect.Width
		h   = rect.Height
		pt1 = image.Pt(x, y)
		pt2 = image.Pt(x+w, y+h)
	)
	for {
		inFile, ok := <-q // closeされるとokがfalseになる
		if !ok {
			return
		}
		base := filepath.Base(inFile)
		outFile := outDir + "/" + base

		src, err := readImageFile(inFile)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		cImg, err := cutter.Crop(src, cutter.Config{
			Width:  w,
			Height: h,
			Anchor: image.Pt(x, y),
			Mode:   cutter.TopLeft,
		})
		if err != nil {
			errs = append(errs, err)
			continue
		}

		dist := image.NewRGBA(image.Rectangle{pt1, pt2})
		draw.Draw(dist, dist.Bounds(), cImg, pt1, draw.Over)

		if err := writeImageFile(outFile, dist); err != nil {
			errs = append(errs, err)
			continue
		}

		createdFiles = append(createdFiles, outFile)
	}
}
