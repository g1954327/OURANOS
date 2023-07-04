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
	//     -t, --token <TOKEN>      specify the token for the service. This option is mandatory.
	//     -h, --help               print this mesasge and exit.
	//     -v, --version            print the version and exit.
	// ARGUMENT
	//     URL     specify the url for shortening. this arguments accept multiple values.
	//             if no arguments were specified, ouranos prints the list of available shorten urls.
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./ouranos", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
