package main

import (
    "flag"
    "fmt"
    "knowledge_base/internal/server"
    "knowledge_base/internal/util"
    "strconv"
)

const serveCommandName = "serve"

const defaultHost = "localhost"
const defaultPort = 8080

type ServeCommand struct {
    flagSet *flag.FlagSet

    host string
    port int

    addr string // TODO

    path string
}

func NewServeCommand() *ServeCommand {
    cmd := &ServeCommand{}

    cmd.flagSet = flag.NewFlagSet(serveCommandName, flag.ExitOnError)

    cmd.flagSet.StringVar(&cmd.host, "host", defaultHost, "host to listen on")
    cmd.flagSet.IntVar(&cmd.port, "port", defaultPort, "port to listen on")

    return cmd
}

func (cmd *ServeCommand) Name() string {
    return serveCommandName
}

func (cmd *ServeCommand) PrintHelp() {
    fmt.Println("Usage sample: ")
    fmt.Println("kbm serve path/to/kb/root")
}

func (cmd *ServeCommand) Execute(args []string) error {
    //if len(args) < 3 {
    //    return errors.New("missing path argument")
    //}
    //
    //var err error
    //
    //cmd.path, err = filepath.Abs(args[2])
    //if err != nil {
    //    return err
    //}

    addr := cmd.host + ":" + strconv.Itoa(cmd.port)

    workingDir, err := util.GetAbsWorkingDir()
    if err != nil {
        return err
    }

    return serve(*workingDir, addr)
}

func serve(kbRootDir string, addr string) error {
    kbServer := server.New(kbRootDir, addr)

    return kbServer.Serve()

    //addr := cmd.host + ":" + strconv.Itoa(cmd.port)
    //
    //serveMux := http.NewServeMux()
    //serveMux.HandleFunc("/", cmd.kbRecords())
    //
    //// TODO http.ServeFile()
    //
    //server := &http.Server{
    //    Addr:    addr,
    //    Handler: serveMux,
    //}
    //return server.ListenAndServe()
}

//func (cmd *ServeCommand) mainPage() http.HandlerFunc {
//    return func(writer http.ResponseWriter, request *http.Request) {
//        // TODO
//    }
//}
//
//func (cmd *ServeCommand) kbRecords() http.HandlerFunc {
//
//    return func(writer http.ResponseWriter, request *http.Request) {
//        rootDir := cmd.path + request.URL.String()
//
//        rootElementPath := rootDir
//        if !strings.HasSuffix(rootElementPath, "/") {
//            rootElementPath += "/"
//        }
//        rootElementPath += "info.json"
//        rootElement, err := parseItemInfo(rootElementPath)
//        if err != nil {
//            fmt.Fprintln(writer, "Failure: ", err)
//            return
//        }
//
//        elements, err := getElements(rootDir)
//        if err != nil {
//            fmt.Fprintln(writer, "Failure: ", err)
//            return
//        }
//
//        var elementContent = []byte("TODO: parse content")
//
//        if rootElement.ElementType == "record" {
//            rootElementContentPath := rootDir
//            if !strings.HasSuffix(rootElementContentPath, "/") {
//                rootElementContentPath += "/"
//            }
//            rootElementContentPath += "content.md"
//            elementContentRaw, err := ioutil.ReadFile(rootElementContentPath)
//            if err != nil {
//                fmt.Fprintln(writer, "Failure: ", err)
//                return
//            }
//            elementContent = blackfriday.Run(elementContentRaw)
//        }
//
//        fmt.Fprintln(writer, "<html>")
//        fmt.Fprintln(writer, "<head>")
//        fmt.Fprintln(writer, "</head>")
//        fmt.Fprintln(writer, "<body>")
//
//        fmt.Fprintf(writer, "<h1>%s</h2>", rootElement.Title)
//        fmt.Fprintf(writer, "<p>%s</p>", rootElement.Description)
//
//        fmt.Fprintln(writer, "<hr/>")
//
//        switch rootElement.ElementType {
//        case "group":
//            for _, item := range elements {
//                // request.URL.String() + "/" +
//                itemUrl := item.Id
//                fmt.Fprintf(writer, "<h2><a href='%s'>%s</a></h2>", itemUrl, item.Title)
//                fmt.Fprintf(writer, "<p>%s</p>", item.Description)
//            }
//        case "record":
//            writer.Write(elementContent)
//        }
//
//        fmt.Fprintln(writer, "</body>")
//        fmt.Fprintln(writer, "</html>")
//    }
//}
//
//func getElements(dir string) ([]*record.Record, error) {
//    elements := make([]*record.Record, 0)
//
//    files, err := ioutil.ReadDir(dir)
//    if err != nil {
//        return nil, err
//    }
//    for _, fileInfo := range files {
//        if !fileInfo.IsDir() {
//            continue
//        }
//        itemInfoPath := dir + "/" + fileInfo.Name() + "/info.json"
//        itemInfo, err := parseItemInfo(itemInfoPath)
//        if err != nil {
//            return nil, err
//        }
//        itemInfo.Id = fileInfo.Name()
//
//        elements = append(elements, itemInfo)
//    }
//
//    return elements, nil
//}
//
//func parseItemInfo(path string) (*record.Record, error) {
//    itemInfo := &record.Record{}
//
//    itemInfoFileBytes, err := ioutil.ReadFile(path)
//    if err != nil {
//        return nil, err
//    }
//
//    err = json.Unmarshal(itemInfoFileBytes, itemInfo)
//    if err != nil {
//        return nil, err
//    }
//
//    return itemInfo, nil
//}
