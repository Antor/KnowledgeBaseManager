package server

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "knowledge_base/internal/record"
    "path/filepath"
)

type RecordPageInfo struct {
    Parents []*record.Record

    Self    *record.Record
    Content template.HTML

    Children []*record.Record
}

func LoadRecordPageInfo(kbRootDirPath string, recordId string) (*RecordPageInfo, error) {
    page := &RecordPageInfo{}

    var err error

    page.Parents, err = loadParents(kbRootDirPath, recordId)
    if err != nil {
        return nil, err
    }

    page.Self, err = record.LoadRecord(kbRootDirPath, recordId)
    if err != nil {
        return nil, err
    }


    page.Content, err = record.LoadRecordContent(kbRootDirPath, recordId)
    if err != nil {
        return nil, err
    }

    page.Children, err = loadChildren(kbRootDirPath, recordId)
    if err != nil {
        return nil, err
    }

    return page, nil
}

func loadParents(kbRootDirPath string, recordId string) ([]*record.Record, error) {
    fmt.Println("loadParents: Begin")

    parents := make([]*record.Record, 0)

    parentId := recordId
    for {
        fmt.Println("loadParent: " + parentId)
        if parentId == "/" {
            break
        }
        parentId, _ = filepath.Split(parentId)
        parentId = filepath.Clean(parentId)


        parent, err := record.LoadRecord(kbRootDirPath, parentId)
        if err != nil {
            return nil, err
        }

        parents = append([]*record.Record{parent}, parents...)

    }

    fmt.Println("loadParents: End")

    return parents, nil
}

func loadChildren(kbRootDirPath string, recordId string) ([]*record.Record, error) {
    recordDirPath := kbRootDirPath + recordId

    childFileInfos, err := ioutil.ReadDir(recordDirPath)
    if err != nil {
        return nil, err
    }

    children := make([]*record.Record, 0)
    for _, childFileInfo := range childFileInfos {
        if !childFileInfo.IsDir() {
            continue
        }

        childRecordId := filepath.Clean(recordId + "/" + childFileInfo.Name())
        childRecord, err := record.LoadRecord(kbRootDirPath, childRecordId)
        if err != nil {
            return nil, err
        }
        children = append(children, childRecord)
    }

    return children, nil
}


