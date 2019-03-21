package domain

// PasteParam はPaste関数の引数に指定するパラメータ
type PasteParam struct {
	Row       int    // 画像タイルのタイル数（行数）
	Col       int    // 画像タイルのタイル数（列数）
	Width     int    // 画像1枚あたりの横幅
	Height    int    // 画像1枚あたりの縦幅
	OutPrefix string // 出力ファイル名のプレフィクス
	NumPadFmt string // 出力ファイル名末尾に付与するナンバリング書式
}
