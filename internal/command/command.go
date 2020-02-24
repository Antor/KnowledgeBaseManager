package command

type Command interface {
    Name() string
    PrintHelp()
    Execute([]string) error
}
