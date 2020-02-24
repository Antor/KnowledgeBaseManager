package record

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/russross/blackfriday/v2"
    "html/template"
    "io/ioutil"
    "os"
    "time"
)

type Record struct {
    Id string `json:"-"`

    CreateDate time.Time `json:"create_date"`

    Title string `json:"title"`

    Description string `json:"description"`
}

const recordFileName = "info.json"
const contentFileName = "content.md"

func LoadRecord(kbRootDirPath string, recordId string) (*Record, error) {
    recordFilePath := recordFilePath(kbRootDirPath, recordId)
    fmt.Println("LoadRecord: " + recordFilePath)

    recordJsonBytes, err := ioutil.ReadFile(recordFilePath)
    if err != nil {
        return nil, err
    }

    record := &Record{}
    record.Id = recordId
    err = json.Unmarshal(recordJsonBytes, record)
    if err != nil {
        return nil, err
    }

    return record, nil
}

func SaveRecord(kbRootDirPath string, record *Record) error {
    recordDir := kbRootDirPath + record.Id
    err := os.Mkdir(recordDir, 0777)
    if err != nil {
        return err
    }

    recordJson, err := json.MarshalIndent(record, "", "    ")
    if err != nil {
        return err
    }

    recordFilePath := recordFilePath(kbRootDirPath, record.Id)
    _, err = os.Stat(recordFilePath)
    if err == nil {
        return errors.New(fmt.Sprintf("file %s already exist", recordFilePath))
    }

    recordFile, err := os.OpenFile(recordFilePath, os.O_CREATE|os.O_WRONLY, 0777)
    if err != nil {
        return err
    }

    _, err = recordFile.Write(recordJson)
    if err != nil {
        _ = recordFile.Close()
        return err
    }

    err = recordFile.Close()
    if err != nil {
        return err
    }

    return nil
}

func recordFilePath(kbRootDirPath string, recordId string) string {
    return kbRootDirPath + recordId + "/" + recordFileName
}

func LoadRecordContent(kbRootDirPath string, recordId string) (template.HTML, error) {
    contentFilePath := contentFilePath(kbRootDirPath, recordId)

    var content template.HTML

    _, err := os.Stat(contentFilePath)
    if err != nil {
        // file not exist
        content = ""
        return content, nil
    }

    contentBytes, err := ioutil.ReadFile(contentFilePath)
    if err != nil {
        return "", err
    }
    content = template.HTML(blackfriday.Run(contentBytes))

    return content, nil
}

func SaveRecordContent(kbRootDirPath string, recordId string, content string) error {
    contentFilePath := contentFilePath(kbRootDirPath, recordId)

    _, err := os.Stat(contentFilePath)
    if err == nil {
        return errors.New(fmt.Sprintf("file %s already exist", contentFilePath))
    }

    contentFile, err := os.OpenFile(contentFilePath, os.O_CREATE|os.O_WRONLY, 0777)
    if err != nil {
        return err
    }

    _, err = contentFile.Write([]byte(content))
    if err != nil {
        _ = contentFile.Close()
        return err
    }

    err = contentFile.Close()
    if err != nil {
        return err
    }

    return nil
}

func contentFilePath(kbRootDirPath string, recordId string) string {
    return kbRootDirPath + recordId + "/" + contentFileName
}


// id of record is string determined by name of directory in which the rest of related content is located

// creation date

// title. human readable

// ------

// root content stored in content.md file. it can reference other records

// meta information stored in meta.json file

// there can be cards for spaced repetition

// there can be related articles from internet (may be store as separate KB records)

// there can be images, videos and audios associated with this KB record

// there can be related books (may be store as separate KB records)

// reference

// Think about JAVA classes, packages and directory structure to store them in context of KB dirs structure

// Initial idea. Directory with KB should have flat structure, all records are represented by dir directly
// within KB directory. Names of dirs are ids of records by which they are referenced.
// record name can contain only latin letters, digits, dots, dashes, underscores.

// record name should be globally unique across all KB directories from which it can be referenced

// There can be multiple KB directories. KB directory can contain only records and no other KBs. No nesting
// This is to simplify reasoning about records

// may be KB records can be of different type: note, event, article, book, etc.
// each record should be self sufficient source of information

//-----------------
// Start by implementing the most simple functionality KB records as notes which can reference other note

// referencing other note should be kind of url or relative path

//-----

// TODO: implement command to capture state of KB record by calculating hash sum of its content and store it in
// meta.json alongside with date of snapshot and comment.
