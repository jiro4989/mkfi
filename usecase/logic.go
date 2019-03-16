package usecase

import (
	"os"
	"runtime"
	"sync"
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
