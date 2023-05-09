package main
import (
    "fmt"

    "github.com/spf13/pflag"
)

type options struct { 
  copy string
  long string
  help bool
  version bool 
}
func buildOptions(args []string) (*options, *flag.FlagSet) { 
  opts := &options{}
  flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
  flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
  flags.StirngVarP(&opts.copy, "copy", "c", "", "出力した短縮URLをコピーする。つまりcmd+vを押すと短縮URLがペーストできる。")
  flags.StirngVarP(&opts.long, "long", "l", "", "短縮する前のURLを表示する")   
  flags.BoolVarP(&opts.help, "help", "h", false, "ヘルプを表示して終了する。")  
  flags.BoolVarP(&opts.version, "version", "v", false, "バージョンを表示して終了する。")
  return opts, flags
}
func perform(opts *options, args []string) *ouranosError { 
  fmt.Println("Hello World")
  return nil
}
func parseOptions(args []string) (*options, []string, *ouranosError) { 
  opts, flags := buildOptions(args)
  flags.Parse(args[1:])
  if opts.help {
    fmt.Println(helpMessage(args[0]))
    return nil, nil, &ouranosError{statusCode: 0, message: ""}
  }
  if opts.copy == "" {
    return nil, nil,
      &ouranosError{statusCode: 3, message: "no url was given"}
  }
  return opts, flags.Args(), nil
}
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
