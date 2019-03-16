package usecase

import (
	"errors"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/disintegration/imaging"
)

func ScaleImageFiles(scaleSize int, outDir string, targetFiles []string) ([]string, error) {
	return logic(scaleSize, outDir, scaleImageFile, targetFiles)
}

func scaleImageFile(wg *sync.WaitGroup, q chan string, in1 interface{}, outDir string, createdFiles []string, errs []error) {
	defer wg.Done()

	scaleSize, ok := in1.(int)
	if !ok {
		err := errors.New(fmt.Sprintf("illegal input data. input=%v", in1))
		errs = append(errs, err)
		return
	}

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

		w := src.Bounds().Size().X * scaleSize / 100
		dist := imaging.Resize(src, w, 0, imaging.Lanczos)
		if err := writeImageFile(outFile, dist); err != nil {
			errs = append(errs, err)
			continue
		}

		if err := writeImageFile(outFile, dist); err != nil {
			errs = append(errs, err)
			continue
		}

		createdFiles = append(createdFiles, outFile)
	}
}
