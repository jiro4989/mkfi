package main

import (
	"net/http"

	"github.com/jiro4989/mkfi/infrastructure"
)

func RootPage(w http.ResponseWriter, r *http.Request, p infrastructure.HTTPParams) {
}

func GenerateChain(w http.ResponseWriter, r *http.Request, p infrastructure.HTTPParams) {
}

// Save はアップロードされた複数の画像ファイルをディレクトリ配下に保存する
func Save(w http.ResponseWriter, r *http.Request, p infrastructure.HTTPParams) {
}

func Generate(w http.ResponseWriter, r *http.Request, p infrastructure.HTTPParams) {
}

func Trim(w http.ResponseWriter, r *http.Request, p infrastructure.HTTPParams) {
}

func Flip(w http.ResponseWriter, r *http.Request, p infrastructure.HTTPParams) {
}

func Paste(w http.ResponseWriter, r *http.Request, p infrastructure.HTTPParams) {
}
