package main

import (
    "errors"
    "flag"
    "fmt"
    "knowledge_base/internal/record"
    "knowledge_base/internal/util"
    "time"
)

const addCommandName string = "add"

type AddCommand struct {
    flagSet *flag.FlagSet
}

func NewAddCommand() *AddCommand {
    cmd := new(AddCommand)
    cmd.flagSet = flag.NewFlagSet(addCommandName, flag.ExitOnError)

    return cmd
}

func (cmd *AddCommand) Name() string {
    return addCommandName
}

func (cmd *AddCommand) PrintHelp() {
    fmt.Println("add command allows to add record or record group to knowledge base")
    fmt.Println("Sample usage: ")
    fmt.Println("kbm add record_name")
    cmd.flagSet.PrintDefaults()
}

func (cmd *AddCommand) Execute(args []string) error {
    err := cmd.flagSet.Parse(args[2:])
    if err != nil {
        return err
    }

    if cmd.flagSet.NArg() == 0 {
        return errors.New("missing record name argument")
    }
    name := cmd.flagSet.Arg(0)

    workingDir, err := util.GetAbsWorkingDir()
    if err != nil {
        return err
    }

    return createRecord(*workingDir, name)
}

func createRecord(workingDir string, name string) error {
    rec := &record.Record{
        Id:          "/" + name,
        CreateDate:  time.Now(),
        Title:       name,
        Description: "Description of " + name,
    }
    err := record.SaveRecord(workingDir, rec)
    if err != nil {
        return err
    }

    recContent := "Initial content for " + name
    err = record.SaveRecordContent(workingDir, rec.Id, recContent)
    if err != nil {
        return err
    }

    return nil
}

//-----------------
//-----------------
//-----------------
//-----------------

//func (cmd *AddCommand) create() error {
//    switch cmd.elementType {
//    case elementTypeRecord:
//        return cmd.createRecord()
//    case elementTypeGroup:
//        return cmd.createGroup()
//    default:
//        return errors.New("unknown element type: " + cmd.elementType)
//    }
//}
//
//func (cmd *AddCommand) createRecord() error {
//    recordDir := cmd.elementPath
//
//    // create record dir
//    err := createDirIfNotExist(recordDir)
//    if err != nil {
//        return err
//    }
//
//    // create info file stub
//    recordInfo := &record.Record{
//        CreateDate:  time.Now(),
//        Title:       "TODO: edit title",
//        Description: "TODO: edit description",
//        ElementType: elementTypeRecord,
//    }
//    err = createInfoFile(recordDir, recordInfo)
//    if err != nil {
//        return err
//    }
//
//    // create content file stub
//    contentFilePath := recordDir + "/content.md"
//    contentFile, err := os.OpenFile(contentFilePath, os.O_CREATE|os.O_WRONLY, 0777)
//    if err != nil {
//        return err
//    }
//    err = contentFile.Close()
//    if err != nil {
//        return err
//    }
//
//    return nil
//}
//
//func (cmd *AddCommand) createGroup() error {
//    groupDir := cmd.elementPath
//
//    err := createDirIfNotExist(groupDir)
//    if err != nil {
//        return err
//    }
//
//    recordInfo := &record.Record{
//        CreateDate:  time.Now(),
//        Title:       "TODO: edit title",
//        Description: "TODO: edit description",
//        ElementType: elementTypeGroup,
//    }
//    err = createInfoFile(groupDir, recordInfo)
//    if err != nil {
//        return err
//    }
//
//    return nil
//}
//
//func createDirIfNotExist(dir string) error {
//    _, err := os.Stat(dir)
//    recordDirExist := err == nil
//    if recordDirExist {
//        files, err := ioutil.ReadDir(dir)
//        if err != nil {
//            return err
//        }
//        if len(files) > 0 {
//            return errors.New("record dir " + dir + " already exist and not empty")
//        }
//
//    } else {
//        err = os.Mkdir(dir, 0777)
//        if err != nil {
//            return err
//        }
//    }
//    return nil
//}
//
//func createInfoFile(parentDir string, info *record.Record) error {
//    infoFilePath := parentDir + "/info.json"
//    infoFile, err := os.OpenFile(infoFilePath, os.O_CREATE|os.O_WRONLY, 0777)
//    if err != nil {
//        return err
//    }
//
//    recordInfoJson, err := json.MarshalIndent(info, "", "    ")
//    if err != nil {
//        _ = infoFile.Close()
//        return err
//    }
//
//    _, err = infoFile.Write(recordInfoJson)
//    if err != nil {
//        _ = infoFile.Close()
//        return err
//    }
//
//    err = infoFile.Close()
//    if err != nil {
//        return err
//    }
//
//    return nil
//}
