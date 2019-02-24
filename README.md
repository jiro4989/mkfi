# mkfi (Make facetile image)

大量の画像ファイルの一括操作をするためのCLIツール兼ローカルサーバアプリ

## 目的

[tkimgutil](https://github.com/jiro4989/tkimgutil)という大量画像処理用のCLIを過去に作った。
このツールはその過去ツールの改良版である。

tkimgutilは以下の機能を持つ

- パイプラインでの大量画像ファイルの一括操作
  - 拡大縮小
  - トリミング
  - 左右反転
  - 結合

しかしながら、以下の課題を抱えていた。

- 拡大、縮小などのスケール値やトリミング位置を実際の画像を確認しながら指定できない

本ツールでは、これらの課題を解決した新規ツールを目指すもの。

## 方式

CLI起動に必要となるパラメータを画像を確認しながら指定できない、という課題について
Webブラウザの画面を介することで解決する。
それ意外の機能については、tkimgutilが保持していた機能はすべて引き継ぐものとする。
（実装ロジックの見直しは行うが）

### UI操作

アプリとしてはバイナリ一つで動くものにする。

画面はlocalhostのWebサーバを起動し、ブラウザの画面に画像をDandDで配置し、
トリミング位置、スケール位置などを指定する。

<!-- TODO
画像を単純に一括拡縮するか、その後の分割もするか、などについても
画面UI上から指定できるものとする。
-->

画像処理を行う機能については、WebAPIとしてやりとりする。
ブラウザ間とlocalhost通信し、最終的に成果物をZIP圧縮して返却する。

### CLI操作

基本的にはtkimgutilと同じサブコマンドとパイプでの処理連続で使うものとする。
違いは、今までサブコマンドが関数ベタ書きだったものが、WebAPIで使用している
関数と共用する点である。

## 使い方

サーバ起動＋画面表示

```bash
mkfi server
```

```bash
mkfi scale -s 90 target.png |
  mkfi trim -x 0 -y 0 --width 144 --height 144 |
  sort |
  mkfi paste -r 4 -c 2
```

## 機能

### paste

RequestParam

```json
{
  "outFileNameFormat":"actor_%03d.png",
  "pattern":[
    ["body1.png", "eye1.png", "mouse1.png"],
    ["body1.png", "eye1.png", "mouse2.png"],
    ["body1.png", "eye2.png", "mouse1.png"],
    ["body1.png", "eye2.png", "mouse2.png"],
  ]
}
```

### trim

RequestParam

```json
{
  "outFileNameFormat":"face_%03d.png",
  "rectangle":{
    "x":0.0,
    "y":0.0,
    "width":144.0,
    "height":144.0
  },
  "pattern":[
    "actor001.png",
    "actor002.png",
    "actor003.png",
    "actor004.png",
    "actor005.png",
    "actor006.png",
    "actor007.png",
    "actor008.png",
  ]
}
```

