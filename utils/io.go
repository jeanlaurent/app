package utils

import (
    "os"
)

// CreateFileWithData creates a new file at the given path and writes
// the provided bytes
func CreateFileWithData(fileName string, data []byte) error {
    f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    defer f.Close()
    _, err = f.Write(data)
    return err
}
