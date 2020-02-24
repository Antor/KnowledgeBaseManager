package main

import "fmt"

type HelpCommand struct {

}

func NewHelpCommand() *HelpCommand {
    return &HelpCommand{};
}

func (cmd *HelpCommand) Name() string {
    return "help"
}

func (cmd *HelpCommand) PrintHelp() {
    fmt.Println("TODO: implement help command help")
}

func (cmd *HelpCommand) Execute([]string) error {
    fmt.Println("Knowledge Base Manager")
    fmt.Println("Version: " + AppVersion)
    fmt.Println("TODO: implement help command")
    return nil
}
