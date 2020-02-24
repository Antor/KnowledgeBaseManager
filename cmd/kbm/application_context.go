package main

import (
    "errors"
    "knowledge_base/internal/command"
)

type ApplicationContext struct {
    commands []command.Command
}

func NewApplicationContext() *ApplicationContext {
    context := &ApplicationContext{}

    context.commands = []command.Command{
        NewHelpCommand(),
        NewAddCommand(),
        NewServeCommand(),
    }

    return context
}

func (context *ApplicationContext) SelectCommand(args []string) (command.Command, error) {
    var commandName string
    if len(args) > 1 {
        commandName = args[1]

    } else {
        commandName = "help"
    }

    var selectedCommand command.Command = nil
    for _, cmd := range context.commands {
        if cmd.Name() == commandName {
            selectedCommand = cmd
        }
    }

    if selectedCommand == nil {
        return nil, errors.New("Unknown command name: " + commandName)
    }

    return selectedCommand, nil
}
