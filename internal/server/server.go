package server

import (
    "fmt"
    "html/template"
    "net/http"
    "os"
    "path/filepath"
)

type KnowledgeBaseServer struct {
    rootDirPath string
    tcpAddr     string
}

func New(rootDirPath string, serverAddr string) *KnowledgeBaseServer {
    server := &KnowledgeBaseServer{
        rootDirPath: rootDirPath,
        tcpAddr:     serverAddr,
    }
    return server
}

func (kbServer *KnowledgeBaseServer) Serve() error {
    server := &http.Server{
        Addr:    kbServer.tcpAddr,
        Handler: kbServer,
    }

    return server.ListenAndServe()
}

func (kbServer *KnowledgeBaseServer) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
    requestedFilePath := kbServer.rootDirPath + request.URL.Path

    requestedFileInfo, err := os.Stat(requestedFilePath)
    if err != nil {
        fmt.Println(err)
        return
    }
    if requestedFileInfo.IsDir() {
        recordId := filepath.Clean(request.URL.Path)
        serveRecord(responseWriter, request, kbServer.rootDirPath, recordId)

    } else {
        http.ServeFile(responseWriter, request, requestedFilePath)
    }
}

func serveRecord(
    responseWriter http.ResponseWriter,
    request *http.Request,
    kbRootDir string,
    recordId string) {

    //_, _ = fmt.Fprint(responseWriter, "TODO: serve record: "+recordId+"\n")

    pageInfo, err := LoadRecordPageInfo(kbRootDir, recordId)
    if err != nil {
        fmt.Println(err)
        return
    }

    executablePath, err := os.Executable()
    if err != nil {
        fmt.Println(err)
        return
    }
    kbmDirPath, _ := filepath.Split(executablePath)

    recordTemplatePath := kbmDirPath + "/web/templates/record_template.html"
    recordTemplate := template.New("record_template.html")
    recordTemplate = recordTemplate.Funcs(
        template.FuncMap{
            "inc": func(i int) int {
                return i + 1
            }})
    recordTemplate, err = recordTemplate.ParseFiles(recordTemplatePath)

    if err != nil {
        fmt.Println(err)
        return
    }

    err = recordTemplate.Execute(responseWriter, pageInfo)
    if err != nil {
        fmt.Println(err)
        return
    }
}
