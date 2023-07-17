package main

import (
	"testing"
)

func Example_Version() {
	goMain([]string{"./ouranos", "--version"})
	// Output:
	// ouranos version 0.2.1
}

func Example_Token() {
	goMain([]string{"./ouranos", "--token"})
	// Output:
	// トークンが与えられていません
}
func Example_Delete() {
	goMain([]string{"./shortURLz", "--delete"})
	// Output:
	// トークンが与えられていません
}

func Example_Completion() {
	goMain([]string{"./ouranos", "--generate-completions"})
	// Output:
	// GenerateCompletion
	// トークンが与えられていません
}

func Example_Help_(){
	goMain([]string{"./ouranos", "--help"})
	//Output:
	// ouranos [OPTIONS] [URLs...]
	//     OPTIONS
	//         -t, --token <TOKEN>      サービスのトークンを指定します。このオプションは必須です。
	//         -h, --help               ヘルプメッセージを表示します。
	//         -v, --version            バージョン情報を表示します。
	//         -p, --past               過去の短縮URLの履歴を5件表示します。
	//         -g, --group <GROUP>      サービスのグループ名を指定します。デフォルトは "ouranos"です。
        //         -d, --delete             指定された短縮URLを削除する。
	//     ARGUMENT
	//         URL                      短縮するURLを指定します。この引数は複数の値を受け付けます。
	//                                  引数が指定されなかった場合、ouranosは利用可能な短縮URLのリストを表示します。
}

func Test_Main_Help(t *testing.T) {
	if status := goMain([]string{"./ouranos", "-h"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}

func Test_Main_Version(t *testing.T) {
	if status := goMain([]string{"./ouranos", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}

func Test_Main_Token(t *testing.T) {
	if status := goMain([]string{"./ouranos", "-t"}); status != 3 {
		t.Error("Expected 3, got ", status)
	}
}

func Test_Main_Past(t *testing.T) {
	if status := goMain([]string{"./ouranos", "-p"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
