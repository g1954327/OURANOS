[![.github/workflows/build.yaml](https://github.com/g1954327/ouranos/actions/workflows/build.yaml/badge.svg)](https://github.com/g1954327/ouranos/actions/workflows/build.yaml)
[![Coverage Status](https://coveralls.io/repos/github/g1954327/ouranos/badge.svg?branch=main)](https://coveralls.io/github/g1954327/ouranos?branch=main)
[![codebeat badge](https://codebeat.co/badges/9f634397-7dff-4ce7-ba2e-d16ed5bce4c2)](https://codebeat.co/projects/github-com-g1954327-ouranos-main)
[![Go Report Card](https://goreportcard.com/badge/github.com/g1954327/ouranos)](https://goreportcard.com/report/github.com/g1954327/ouranos)

<img alt="Github" src="https://img.shields.io/badge/Developer-Shingo_Morikawa-blueviolet"> <img alt="GitHub" src="https://img.shields.io/github/license/g1954327/ouranos"> <img alt="GitHub" src=https://img.shields.io/badge/Langage-GO-blue> <img alt="GitHub" src="https://img.shields.io/badge/Version-0.2.1-important">


# ouranos
URLを渡すと短縮されたURLを出力してくれるCLIです。
## 説明
URLとはインターネット上の所在を表記するものであり、インターネット上のサイトにアクセスするには必ず必要となる。URLは基本的に長くなる。そのためTwitterなどの文字制限があるサイトにURLを掲載する際には、文字数を減らしたい。またURLをQRコードする際に、URLが長いとQRコードに不備が出てしまうかもしれない。これらを解決するためにはURLを機能そのままで短縮したい。そのためのアプリです。CLIで動作させることによって、入力されたURLを短縮URLに変換してくれるアプリです。

## 使用方法
    $ ouranos -h  
    ouranos version : 0.2.1  最終更新 7/14
    ouranos [command] <URL>
    command
        -t, --token <TOKEN>      サービスのトークンを指定します。このオプションは必須です。
        -h, --help               ヘルプメッセージを表示します。
        -v, --version            バージョン情報を表示します。
        -p, --past               過去の短縮URLの履歴を5件表示します。
        -g, --group <GROUP>      サービスのグループ名を指定します。デフォルトは "ouranos"です。
        -d, --delete             指定された短縮URLを削除する。
        
## インストール方法 
    $ brew install g1954327/tap/ouranos
    
# このプロジェクトについて
    
        
## 開発者
京都産業大学大学院 先端情報学研究科 森川 真伍

<img alt="Github" src="https://img.shields.io/badge/Developer-Shingo_Morikawa-blueviolet">

## 使用言語
Go言語

<img alt="GitHub" src=https://img.shields.io/badge/Langage-GO-blue>

## ライセンス
MITライセンス

<img alt="GitHub" src="https://img.shields.io/github/license/g1954327/ouranos">

## バージョン

<img alt="GitHub" src="https://img.shields.io/badge/Version-0.2.1-important">

## ロゴ

![Sickle2](https://user-images.githubusercontent.com/77278892/235562689-71e6140f-12f5-4f92-8c92-b05d7de244d5.png)

## 名前の由来

URLを短くする　→ URLShorter → URLS → ウラルス → ウラノスという発想（ほとんどこじつけ）から。
またウラノスはギリシャ神話の神様であり、あるところを鎌で切られたことから、URLを短くする(つまり切り取る)という発想。
そのため、ロゴも上記の通り鎌とした。

