package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

const outDir = "out"

func RootPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	fmt.Fprintf(w, "%s", `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title></title>
</head>
<body>
    <form method="post" action="/save" enctype="multipart/form-data">
        <fieldset>
            <input type="file" name="upload_files" id="upload_files" multiple="multiple">
            <input type="submit" name="submit" value="アップロード開始">
        </fieldset>
    </form>
</body>
</html>
	`)
}
func GenerateChain(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	// (save) 一旦処理対象の画像ファイルをすべてローカルに保存
	// (generate) 画像をすべて組み合わせてファイル出力。戻り値は生成されたファイル名配列
	// (trim) 画像をすべて指定位置でトリミングしてファイル出力。戻り値は生成されたファイル名配列
	// (flip) 画像をすべて左右反転してファイル出力。戻り値は生成されたファイル名配列
	// (paste) 画像をすべてタイル状に貼り付けてファイル出力。戻り値は生成されたファイル名配列
	// (archive) 一連の処理で生成された成果物を圧縮
	// 圧縮ファイルを返却
}

// Save はアップロードされた複数の画像ファイルをディレクトリ配下に保存する
func Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := save(w, r, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// もとのページにリダイレクト
	http.Redirect(w, r, "/", http.StatusFound)
}

func save(w http.ResponseWriter, r *http.Request, p httprouter.Params) ([]string, error) {
	defer r.Body.Close()

	// multipartリーダーの取得
	mr, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}

	// 保存ファイルの格納先ディレクトリの作成
	subDirName := "save"
	saveDir := outDir + "/" + subDirName
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return nil, err
	}

	var ret []string
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}

		// ファイル名がない場合はスキップ
		if part.FileName() == "" {
			continue
		}

		// defer closeするために無名関数呼び出し
		fn, err := func() (string, error) {
			// 保存ファイルの生成
			saveFilePath := saveDir + "/" + part.FileName()
			saveFile, err := os.Create(saveFilePath)
			if err != nil {
				return "", err
			}
			defer saveFile.Close()

			// ファイルの内容を保存ファイルに書き込み
			_, err = io.Copy(saveFile, part)
			if err != nil {
				return "", err
			}

			return saveFilePath, nil
		}()
		if err != nil {
			return nil, err
		}

		ret = append(ret, fn)
	}

	return ret, nil
}

func Generate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	// TODO
}

func Trim(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()

	// リクエストボディを読み取る
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// リクエストボディの読み取りに失敗した => 400 Bad Requestエラー
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// JSONパラメーターを構造体にする為の定義
	type ExampleParameter struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	var param ExampleParameter

	// ExampleParameter構造体に変換
	err = json.Unmarshal(bodyBytes, &param)
	if err != nil {
		// JSONパラメーターを構造体への変換に失敗した => 400 Bad Requestエラー
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 構造体に変換したExampleParameterを文字列にしてレスポンスに書き込む
	fmt.Fprintf(w, fmt.Sprintf("%+v\n", param))
}

func Flip(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	// TODO
}

func Paste(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	// TODO
}
