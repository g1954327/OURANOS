package main

import (
	"fmt"
	"os"
	"path/filepath"
	flag "github.com/spf13/pflag"
)

const VERSION = "0.1.1"

//コマンドラインオプションを管理するためのoptions構造体を定義。
type options struct {
	help      bool
	version   bool
	past      bool
	token     string
}

//ouranosError構造体を定義し、エラーメッセージを表現するためのErrorメソッドの実装。  
type ouranosError struct {
	statusCode int
	message    string
}

func bitlyRequest(opts *options, long_url *string) {
	fmt.Printf("long_url: %s\n", *long_url)
}

var completions bool

//オプションの定義とオプションを解析するためのbuildOptions関数を定義。 
func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := &options{}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.StringVarP(&opts.token, "token", "t", "", "サービスのトークンを指定します。このオプションは必須です。")
	flags.BoolVarP(&opts.past, "past", "p", false, "過去の履歴を5件表示します")
	flags.BoolVarP(&opts.help, "help", "h", false, "ヘルプメッセージを表示します。")
	flags.BoolVarP(&opts.version, "version", "v", false, "バージョンを表示します。")
	flags.BoolVarP(&completions, "generate-completions", "", false, "generate completions") 
	flags.MarkHidden("generate-completions")
	return opts, flags
}
//オプションと引数をもとに実行する操作を決定し、実行するためのperform関数を定義。
func perform(opts *options, args []string) *ouranosError {
	if opts != nil {
		fmt.Printf("Token: %s\n", opts.token)
	}
	for _, long_url := range args {
		bitlyRequest(opts, &long_url)
	}
	return nil
}

//オプションを解析し、options構造体と引数を取得するためのparseOptions関数を定義。

func parseOptions(args []string) (*options, []string, *ouranosError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	f, err := os.OpenFile("past.txt", os.O_RDWR|os.O_APPEND,0666)
	_, err = f.WriteString("go run cmd/main.go ")
	_, err = f.WriteString(args[1])
	_, err = f.WriteString("\n")
	//fmt.Println("go run cmd/main.go",args[1:])
	if err != nil {
		fmt.Println("fail to read file")
	}
	defer f.Close()
	if opts.help {
		fmt.Println(helpMessage(args))
		_, err = f.WriteString(helpMessage(args))
		_, err = f.WriteString("\n")
		_, err = f.WriteString("\n")
		return nil, nil, &ouranosError{statusCode: 0, message: ""}
	}
	if opts.version {
		fmt.Println(versionString(args))
		_, err = f.WriteString(versionString(args))
		_, err = f.WriteString("\n")
		_, err = f.WriteString("\n")
		return nil, nil, &ouranosError{statusCode: 0, message: ""}
	}
	if opts.past {
		fmt.Println(pastString(args))
		_, err = f.WriteString(pastString(args))
		_, err = f.WriteString("\n")
		_, err = f.WriteString("\n")
		return nil, nil, &ouranosError{statusCode: 0, message: ""}
	}
	if opts.token == "" {
		return nil, nil, &ouranosError{statusCode: 3, message: "トークンが与えられていません"}
	}
	return opts, flags.Args(), nil
}

//ヘルプメッセージの出力
func helpMessage(args []string) string {
	prog := "ouranos"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s [OPTIONS] [URLs...]
    OPTIONS
        -t, --token <TOKEN>      サービスのトークンを指定します。このオプションは必須です。
        -h, --help               ヘルプメッセージを表示します。
        -v, --version            バージョン情報を表示します。
        -p, --past               過去の履歴を5件表示します。
    ARGUMENT
        URL                      短縮するURLを指定します。この引数は複数の値を受け付けます。
                                 引数が指定されなかった場合、ouranosは利用可能な短縮URLのリストを表示します。`, prog)
}
//バージョン情報の出力
func versionString(args []string) string {
	prog := "ouranos"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version %s", prog, VERSION)
}

func pastString(args []string) string {
	f, err := os.OpenFile("past.txt", os.O_RDWR|os.O_APPEND,0666)
	//fmt.Println("go run main.go",args[1])
	if err != nil {
		fmt.Println("fail to read file")
	}
	defer f.Close()

	f2, err2 := os.OpenFile("past.txt", os.O_RDWR|os.O_APPEND,0666)
	data := make([]byte, 1024)
	count, err2 := f2.Read(data)
	if err2 != nil {
		fmt.Println("fail to read file")
	}
	fmt.Println(string(data[:count]))
	defer f2.Close()

	return fmt.Sprintf("past %s", VERSION)
}

func (e ouranosError) Error() string {
	return e.message
}

//メイン関数の実装(goMain関数)を定義し、parseOptionsとperformを呼び出す。 
func goMain(args []string) int {
	opts, args, err := parseOptions(args)
	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	if err := perform(opts, args); err != nil {
		fmt.Println(err.Error())
		return err.statusCode
	}
	return 0
}
//メイン関数での処理結果に応じて、終了ステータスを返す。
func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
