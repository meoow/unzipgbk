package main

import (
    "archive/zip"
    "fmt"
    "io"
    "log"
    "os"
    "sort"
    "strings"

    iconv "github.com/sloonz/go-iconv"
)

type SortFile []*zip.File

func init() {
    log.SetOutput(os.Stderr)
}

func main() {
    rzip, err := zip.OpenReader(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer rzip.Close()
    sort.Sort(SortFile(rzip.File))
    for _, file := range rzip.File {
        utf8name, _ := iconv.Conv(file.Name, "utf-8", "gbk")
        if strings.HasSuffix(file.Name, "/") {
            os.MkdirAll(utf8name, 0755)
        } else {
            err := extractZip(utf8name, file)
            if err != nil {
                log.Println(err)
            } else {
                fmt.Println(utf8name)
            }
        }
    }
}

func extractZip(dst string, zf *zip.File) error {
    zfReader, err := zf.Open()
    if err != nil {
        return err
    }
    defer zfReader.Close()

    flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
    dstWriter, err := os.OpenFile(dst, flag, 0644)
    if err != nil {
        return err
    }
    defer dstWriter.Close()

    copiedSize, err := io.Copy(dstWriter, zfReader)
    if err != nil {
        return err
    }
    if uint64(copiedSize) != zf.UncompressedSize64 {
        return fmt.Errorf("Failed to extract file %s: size mismatched.", dst)
    }
    return nil
}

func (f SortFile) Len() int {
    return len(f)
}

func (f SortFile) Swap(i, j int) {
    f[i], f[j] = f[j], f[i]
}

func (f SortFile) Less(i, j int) bool {
    if strings.HasSuffix(f[i].Name, "/") {
        if strings.HasSuffix(f[j].Name, "/") {
            return len(f[i].Name) < len(f[j].Name)
        } else {
            return true
        }
    } else {
        if strings.HasSuffix(f[j].Name, "/") {
            return false
        } else {
            return len(f[i].Name) < len(f[j].Name)
        }
    }
}
