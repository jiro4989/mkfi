package usecase

import (
	"image"
	"image/draw"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/jiro4989/mkfi/domain"
	"github.com/oliamb/cutter"
)

func TrimImageFiles(rect domain.Rectangle, outDir string, targetFiles []string) ([]string, error) {
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	// 処理対象の数だけキューを生成
	q := make(chan string, len(targetFiles))

	// 処理結果を格納用
	var createdFiles []string
	var errs []error

	// ワーカースレッドの生成
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go trimImageFile(&wg, q, rect, outDir, createdFiles, errs)
	}

	// 処理対象ファイル名を送信
	// ワーカースレッドの数分しか並列に処理しない
	for _, f := range targetFiles {
		q <- f
	}
	close(q)
	wg.Wait()

	if 0 < len(errs) {
		return nil, errs[0]
	}

	return createdFiles, nil
}

func trimImageFile(wg *sync.WaitGroup, q chan string, rect domain.Rectangle, outDir string, createdFiles []string, errs []error) {
	var (
		x   = rect.X
		y   = rect.Y
		w   = rect.Width
		h   = rect.Height
		pt1 = image.Pt(x, y)
		pt2 = image.Pt(x+w, y+h)
	)
	defer wg.Done()
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
