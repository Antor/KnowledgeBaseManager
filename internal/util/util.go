package util

import (
    "os"
    "path/filepath"
)

func GetAbsWorkingDir() (*string, error) {
    workingDir, err := os.Getwd()
    if err != nil {
        return nil, err
    }

    workingDir, err = filepath.Abs(workingDir)
    if err != nil {
        return nil, err
    }

    return &workingDir, nil
}
