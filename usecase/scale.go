package usecase

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/disintegration/imaging"
)

func logic(in1 interface{}, outDir string, f func(wg *sync.WaitGroup, q chan string, in1 interface{}, outDir string, createdFiles []string, errs []error), targetFiles []string) ([]string, error) {
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
		go f(&wg, q, in1, outDir, createdFiles, errs)
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
