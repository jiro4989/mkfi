package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GenerateChain(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	// (save) 一旦処理対象の画像ファイルをすべてローカルに保存
	// (generate) 画像をすべて組み合わせてファイル出力。戻り値は生成されたファイル名配列
	// (trim) 画像をすべて指定位置でトリミングしてファイル出力。戻り値は生成されたファイル名配列
	// (flip) 画像をすべて左右反転してファイル出力。戻り値は生成されたファイル名配列
	// (paste) 画像をすべてタイル状に貼り付けてファイル出力。戻り値は生成されたファイル名配列
}

func Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	// TODO
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
