package main

import (
	"testing"
)

func Example_Version() {
	goMain([]string{"./ouranos", "--version"})
	// Output:
	// トークンが指定されていません
}

func Example_Token() {
	goMain([]string{"./ouranos", "--token"})
	// Output:
	// トークンが指定されていません
}

func Example_Past() {
	goMain([]string{"./ouranos", "-past"})
	// Output:
	// トークンが指定されていません
}

func Example_Completion() {
	goMain([]string{"./ouranos", "--generate-completions"})
	// Output:
	// トークンが指定されていません
}

func Example_Help() {
	goMain([]string{"./ouranos", "--help"})
	// Output:
	// ouranos [OPTIONS] [URLs...]
	// OPTIONS
	//     -t, --token <TOKEN>      サービスのトークンを指定します。このオプションは必須です。
	//     -h, --help               ヘルプメッセージを表示します。
	//     -v, --version            バージョン情報を表示します。
	//     -p, --past               過去の履歴を5件表示します。
	// ARGUMENT
	//     URL                      短縮するURLを指定します。この引数は複数の値を受け付けます。
	//                              引数が指定されなかった場合、ouranos は利用可能な短縮 URL のリストを表示します。
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./ouranos", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
