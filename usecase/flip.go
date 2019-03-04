package usecase

import (
	"image"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/disintegration/gift"
)

func FlipImageFiles(outDir string, targetFiles []string) ([]string, error) {
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
		go flipImage(&wg, q, outDir, createdFiles, errs)
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

func flipImage(wg *sync.WaitGroup, q chan string, outDir string, createdFiles []string, errs []error) {
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
