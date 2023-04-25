[![.github/workflows/build.yaml](https://github.com/g1954327/ouranos/actions/workflows/build.yaml/badge.svg)](https://github.com/g1954327/ouranos/actions/workflows/build.yaml)
[![Coverage Status](https://coveralls.io/repos/github/g1954327/ouranos/badge.svg?branch=main)](https://coveralls.io/github/g1954327/ouranos?branch=main)
[![codebeat badge](https://codebeat.co/badges/9f634397-7dff-4ce7-ba2e-d16ed5bce4c2)](https://codebeat.co/projects/github-com-g1954327-ouranos-main)
[![Go Report Card](https://goreportcard.com/badge/github.com/g1954327/ouranos)](https://goreportcard.com/report/github.com/g1954327/ouranos)



# ouranos
URLを渡すと短縮されたURLを出力してくれるCLIです。
## 説明
URLとはインターネット上の所在を表記するものであり、インターネット上のサイトにアクセスするには必ず必要となる。URLは基本的に長くなる。そのためTwitterなどの文字制限があるサイトにURLを掲載する際には、文字数を減らしたい。またURLをQRコードする際に、URLが長いとQRコードに不備が出てしまうかもしれない。これらを解決するためにはURLを機能そのままで短縮したい。そのためのアプリです。CLIで動作させることによって、入力されたURLを短縮URLに変換してくれるアプリです。

## 使用方法
    $ ouranos -h  
    ouranos version : 0.00  最終更新 4/25
    ouranos [command] <URL>
    command
        -v , --version       このアプリケーションのバージョンを説明する。 
        -h , --help          このメッセージを表示する。
        -l , --long          元々のURLも出力する
        -c , --copy          出力した短縮URLをコピーする。つまりcmd+vを押すと短縮URLがペーストできる。
        
## インストール方法 (まだ途中)
    $ brew install ouranos ?
        
## 開発者
    京都産業大学大学院 先端情報学専攻科 森川 真伍
## ロゴ

## 名前の由来
