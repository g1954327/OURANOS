package main

import (
	"fmt"
	"bufio"
	"os"
	"path/filepath"
	"github.com/g1954327/ouranos"
	flag "github.com/spf13/pflag"
)

const VERSION = "0.2.1"

//バージョン情報の出力
func versionString(args []string) string {
	prog := "ouranos"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version %s", prog, VERSION)
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
        -p, --past               過去の短縮URLの履歴を5件表示します。
	-g, --group <GROUP>      サービスのグループ名を指定します。デフォルトは "ouranos"です。
    	-d, --delete             指定された短縮URLを削除する。
    ARGUMENT
        URL                      短縮するURLを指定します。この引数は複数の値を受け付けます。
                                 引数が指定されなかった場合、ouranosは利用可能な短縮URLのリストを表示します。`, prog)
}

//ouranosError構造体を定義し、エラーメッセージを表現するためのErrorメソッドの実装。  
type OuranosError struct {
	statusCode int
	message    string
}
//エラーメッセージを表現するためのメソッド
func (e OuranosError) Error() string {
	return e.message
}

//コマンドラインオプションを管理するためのoptions構造体を定義。
type options struct {
	help      bool
	version   bool
	past      bool
	token     string
	runOpt    *runOpts
	flagSet   *flags
}
//コマンドラインフラグを管理するための構造体
type flags struct {
	deleteFlag     bool
	listGroupFlag  bool
}
//実行オプションを管理するための構造体
type runOpts struct {
	token  string
	config string
	group  string
}

var completions bool
//options 構造体のインスタンスを生成して返すための関数
func newOptions() *options {
	return &options{runOpt: &runOpts{},flagSet: &flags{}}
}

func (opts *options) mode(args []string) ouranos.Mode {
	switch {
	case opts.flagSet.listGroupFlag:
		return ouranos.ListGroup
	case len(args) == 0:
		return ouranos.List
	case opts.flagSet.deleteFlag:
		return ouranos.Delete
	default:
		return ouranos.Shorten
	}
}

//オプションの定義とオプションを解析するためのbuildOptions関数を定義。 
func buildOptions(args []string) (*options, *flag.FlagSet) {
	//opts := &options{}
	opts := newOptions()
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.StringVarP(&opts.runOpt.token, "token", "t", "", "サービスのトークンを指定します。このオプションは必須です。")
	flags.StringVarP(&opts.runOpt.group, "group", "g", "", "サービスのグループ名を指定します。デフォルトは ouranos")
	flags.BoolVarP(&opts.flagSet.listGroupFlag, "list-group", "L", false, "グループをリストアップする。これは隠しオプションです。")
	flags.BoolVarP(&opts.flagSet.deleteFlag, "delete", "d", false, "指定された短縮URLを削除する。")
	flags.BoolVarP(&opts.past, "past", "p", false, "過去の短縮URLの履歴を5件表示します")
	flags.BoolVarP(&opts.help, "help", "h", false, "ヘルプメッセージを表示します。")
	flags.BoolVarP(&opts.version, "version", "v", false, "バージョンを表示します。")
	flags.BoolVarP(&completions, "generate-completions", "", false, "generate completions") 
	flags.MarkHidden("generate-completions")
	return opts, flags
}

//オプションを解析し、options構造体と引数を取得するためのparseOptions関数を定義。
func parseOptions(args []string) (*options, []string, *OuranosError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	//if completions{
	//	fmt.Println("GenerateCompletion")
	//	GenerateCompletion(flags)
	//}
	if opts.help {
		fmt.Println(helpMessage(args))
		return nil, nil, &OuranosError{statusCode: 0, message: ""}
	}
	if opts.version {
		fmt.Println(versionString(args))
		return nil, nil, &OuranosError{statusCode: 0, message: ""}
	}
	if opts.past {
		fmt.Println(pastString(args))
		return nil, nil, &OuranosError{statusCode: 0, message: ""}
	}
	if opts.runOpt.token == "" {
		return nil, nil, &OuranosError{statusCode: 3, message: "トークンが与えられていません"}
	}
	return opts, flags.Args(), nil
}
//過去の短縮URLの履歴を読み込み、最新の5件を表示するための関数
func pastString(args []string) string {
	//ファイルを読み取り開ける
	f, err := os.OpenFile("past.txt", os.O_RDWR|os.O_APPEND,0666)
	if err != nil {
		fmt.Println("fail to read file")
	}
	defer f.Close()

	//ファイルから読み出し
	f2, err2 := os.OpenFile("past.txt", os.O_RDWR|os.O_APPEND,0666)
	data := make([]byte, 1024)
	count, err2 := f2.Read(data)
	if err2 != nil {
		fmt.Println("fail to read file")
	}
	defer f2.Close()

	return string(data[:count])
}
//指定されたURLを短縮し、結果を表示するための関数
func shortenEach(bitly *ouranos.Bitly, config *ouranos.Config, url string) error {
	f, err := os.OpenFile("past.txt", os.O_RDWR|os.O_APPEND,0666)
	line :=0
	var slice []string
	var slice1 []string
	if err != nil {
		fmt.Println("fail to read file")
	}
	scanner := bufio.NewScanner(f)

    //データを１行読み込み
    for scanner.Scan() {
		entry := scanner.Text() //一行を保持
		slice = append(slice,entry)
		line++;
    }
	defer f.Close()

	if line > 8 { //8行より大きかったら(履歴が5件以上あったら)
		f_o, err := os.OpenFile("past.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
  		if err != nil {
    		fmt.Println(err)
  		}
		slice1 = append(slice[:0],slice[line-8:]...)
		for i := 0; i<8; i++{
			_, err = f_o.WriteString(slice1[i])
			_, err = f_o.WriteString("\n")
		}
  		defer f_o.Close()
	}
	result, err := bitly.Shorten(config, url)
	if err != nil {
		return err
	}
	fmt.Println(result)
	_, err = f.WriteString(result.String())
    _, err = f.WriteString("\n")
	_, err = f.WriteString("\n")
	return nil
}
//指定されたURLを削除するための関数
func deleteEach(bitly *ouranos.Bitly, config *ouranos.Config, url string) error {
	return bitly.Delete(config, url)
}
//現在の短縮URLのリストを表示するための関数
func listUrls(bitly *ouranos.Bitly, config *ouranos.Config) error {
	urls, err := bitly.List(config)
	if err != nil {
		return err
	}
	for _, url := range urls {
		fmt.Println(url)
	}
	return nil
}
//グループのリストを表示するための関数
func listGroups(bitly *ouranos.Bitly, config *ouranos.Config) error {
	groups, err := bitly.Groups(config)
	if err != nil {
		return err
	}
	for i, group := range groups {
		fmt.Printf("GUID[%d] %s\n", i, group.Guid)
	}
	return nil
}
//引数のURLごとに指定された関数を実行するための関数
func performImpl(args []string, executor func(url string) error) *OuranosError {
	for _, url := range args {
		err := executor(url)
		if err != nil {
			return makeError(err, 3)
		}
	}
	return nil
}

//オプションと引数をもとに実行する操作を決定し、実行するためのperform関数を定義。
func perform(opts *options, args []string) *OuranosError {
	bitly := ouranos.NewBitly(opts.runOpt.group)
	config := ouranos.NewConfig(opts.runOpt.config, opts.mode(args))
	config.Token = opts.runOpt.token
	switch config.RunMode {
	case ouranos.List:
		err := listUrls(bitly, config)
		return makeError(err, 1)
	case ouranos.ListGroup:
		err := listGroups(bitly, config)
		return makeError(err, 2)
	case ouranos.Delete:
		return performImpl(args, func(url string) error {
			return deleteEach(bitly, config, url)
		})
	case ouranos.Shorten:
		return performImpl(args, func(url string) error {
			return shortenEach(bitly, config, url)
		})
	}
	return nil
}
//エラーを OuranosError 構造体に変換するための関数
func makeError(err error, status int) *OuranosError {
	if err == nil {
		return nil
	}
	ue, ok := err.(*OuranosError)
	if ok {
		return ue
	}
	return &OuranosError{statusCode: status, message: err.Error()}
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
