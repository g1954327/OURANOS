package main

import (
	//"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/g1954327/ouranos"
	flag "github.com/spf13/pflag"
)

const VERSION = "0.1.1"

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
        -p, --past               過去の履歴を5件表示します。
    ARGUMENT
        URL                      短縮するURLを指定します。この引数は複数の値を受け付けます。
                                 引数が指定されなかった場合、ouranosは利用可能な短縮URLのリストを表示します。`, prog)
}

//ouranosError構造体を定義し、エラーメッセージを表現するためのErrorメソッドの実装。  
type OuranosError struct {
	statusCode int
	message    string
}

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

type runOpts struct {
	token  string
	qrcode string
	config string
	group  string
}

var completions bool

func newOptions() *options {
	return &options{runOpt: &runOpts{}, flagSet: &flags{}}
}

func (opts *options) mode(args []string) ouranos.Mode {
	switch {
	case opts.flagSet.listGroupFlag:
		return ouranos.ListGroup
	case len(args) == 0:
		return ouranos.List
	case opts.flagSet.deleteFlag:
		return ouranos.Delete
	case opts.runOpt.qrcode != "":
		return ouranos.QRCode
	default:
		return ouranos.Shorten
	}
}

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

//オプションを解析し、options構造体と引数を取得するためのparseOptions関数を定義。
func parseOptions(args []string) (*options, []string, *OuranosError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	f, err := os.OpenFile("past.txt", os.O_RDWR|os.O_APPEND,0666)
	_, err = f.WriteString("go run cmd/main.go ")
	_, err = f.WriteString(args[1])
	_, err = f.WriteString("\n")
	if err != nil {
		fmt.Println("fail to read file")
	}
	defer f.Close()
	if opts.help {
		fmt.Println(helpMessage(args))
		_, err = f.WriteString(helpMessage(args))
		_, err = f.WriteString("\n")
		_, err = f.WriteString("\n")
		return nil, nil, &OuranosError{statusCode: 0, message: ""}
	}
	if opts.version {
		fmt.Println(versionString(args))
		_, err = f.WriteString(versionString(args))
		_, err = f.WriteString("\n")
		_, err = f.WriteString("\n")
		return nil, nil, &OuranosError{statusCode: 0, message: ""}
	}
	if opts.past {
		fmt.Println(pastString(args))
		_, err = f.WriteString(pastString(args))
		_, err = f.WriteString("\n")
		_, err = f.WriteString("\n")
		return nil, nil, &OuranosError{statusCode: 0, message: ""}
	}
	if opts.token == "" {
		return nil, nil, &OuranosError{statusCode: 3, message: "トークンが与えられていません"}
	}
	return opts, flags.Args(), nil
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

func shortenEach(bitly *ouranos.Bitly, config *ouranos.Config, url string) error {
	result, err := bitly.Shorten(config, url)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func deleteEach(bitly *ouranos.Bitly, config *ouranos.Config, url string) error {
	return bitly.Delete(config, url)
}

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
