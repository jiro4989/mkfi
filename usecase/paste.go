package usecase

import (
	"fmt"
	"image"
	"image/draw"
	"math"
	"os"

	"github.com/jiro4989/mkfi/domain"
	"github.com/jiro4989/mkfi/log"
)

func PasteImageFiles(param domain.PasteParam, outDir string, targetFiles []string) ([]string, error) {
	var (
		row       = param.Row
		col       = param.Col
		width     = param.Width
		height    = param.Height
		outPre    = param.OutPrefix
		padFmt    = param.NumPadFmt
		max       = row * col
		fnFmt     = fmt.Sprintf("%s/%s%s.png", outDir, outPre, padFmt)
		outWidth  = width * col
		outHeight = height * row
		fcnt      = 1 // タイル状に貼り付けられたことで生成された画像ファイルの枚数
	)
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return nil, err
	}

	dist := image.NewRGBA(image.Rect(0, 0, outWidth, outHeight))

	// 処理結果を格納用
	var createdFiles []string
	var errs []error

	for cnt, inFile := range targetFiles {
		src, err := readImageFile(inFile)
		if err != nil {
			errs = append(errs, err)
			log.Error(fmt.Sprintf("Failed reading image file. file=%s, err=%v", inFile, err))
		}

		pt := calcPos(cnt, row, col, width, height, max)

		// 画像を貼り付け
		rect := image.Rectangle{
			pt,
			image.Pt(pt.X+width, pt.Y+height),
		}
		draw.Draw(dist, rect, src, image.Pt(0, 0), draw.Over)

		// 画像の保存
		// 生成するタイル画像内に貼り付けた画像の数が上限に達していた場合は
		// 画像ファイルを出力し、画像バッファを新規作成する。
		if (cnt+1)%max == 0 {
			on := fmt.Sprintf(fnFmt, fcnt)
			if err := writeImageFile(on, dist); err != nil {
				errs = append(errs, err)
				log.Error(fmt.Sprintf("Failed reading image file. file=%s, err=%v", on, err))
			}

			dist = image.NewRGBA(image.Rect(0, 0, outWidth, outHeight))
			fcnt++

			fmt.Println(on)
			createdFiles = append(createdFiles, on)
		}
	}

	// 空のファイルが生成されないようにチェック
	if 0 < len(targetFiles)-1%max {
		on := fmt.Sprintf(fnFmt, fcnt)
		if err := writeImageFile(on, dist); err != nil {
			errs = append(errs, err)
			log.Error(fmt.Sprintf("Failed reading image file. file=%s, err=%v", on, err))
		}
		createdFiles = append(createdFiles, on)
	}

	if 0 < len(errs) {
		return nil, errs[0]
	}

	return createdFiles, nil
}

// calcPos は画像の貼り付け座標を計算して返す。
// 貼り付ける位置は下記のようなタイル状になる
//  1 2 3 4
//  5 6 7 8
// - i index   何枚目の画像を処理しているか
// - r row     タイル画像は何行までのタイルにするか
// - c col     タイル画像は何列までのタイルにするか
// - w width   1タイルあたりの横幅
// - h height  1タイルあたりの縦幅
// - m max     この画像ファイルの中に何枚までタイルを配置できるか
func calcPos(i, r, c, w, h, m int) image.Point {
	var (
		ti = float64(i)
		tc = float64(c)
		tw = float64(w)
		th = float64(h)
		tm = float64(m)
	)

	if m <= i {
		a, _ := math.Modf(ti / tm)
		ti -= a * tm
	}

	x := math.Mod(ti, tc) * tw
	y, _ := math.Modf(ti / tc)
	y *= th

	return image.Pt(int(x), int(y))
}
