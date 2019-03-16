package usecase

import (
	"image"
	"path/filepath"
	"sync"

	"github.com/disintegration/gift"
)

func FlipImageFiles(outDir string, targetFiles []string) ([]string, error) {

	return logic(nil, outDir, flipImageFile, targetFiles)
}

func flipImageFile(wg *sync.WaitGroup, q chan string, _ interface{}, outDir string, createdFiles []string, errs []error) {
	defer wg.Done()

	for {
		inFile, ok := <-q
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

		g := gift.New(gift.FlipHorizontal())
		dist := image.NewRGBA(g.Bounds(src.Bounds()))
		g.Draw(dist, src)

		if err := writeImageFile(outFile, dist); err != nil {
			errs = append(errs, err)
			continue
		}

		createdFiles = append(createdFiles, outFile)
	}
}
