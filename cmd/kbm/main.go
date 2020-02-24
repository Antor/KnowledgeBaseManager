package main

import (
    "fmt"
    "os"
)

const AppVersion = "0.1.0"

func main() {
    context := NewApplicationContext()

    selectedCommand, err := context.SelectCommand(os.Args)

    if err != nil {
        fmt.Println(err)
        return
    }

    err = selectedCommand.Execute(os.Args)
    if err != nil {
        fmt.Println(err)
        return
    }

    // TODO: kbm init

    // TODO: kbm add record <record_name>
    // 1. create file record_name.json pre-filled with meta information
    // 2. create folder record_name with file index.md

    // TODO: kbm add group <group_name>

    // TODO: kbm serve <path_to_kb>

    // TODO: kbm help
    // 1. print list of available commands

    // TODO: kbm help <command>
    // 1. print help for specified command
}
