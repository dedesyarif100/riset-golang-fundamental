package main

import (
    "crypto/md5"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "sync"
    "time"
)

var tempPathThree = filepath.Join(os.Getenv("TEMP"), "FILE TEMPORARY")

type FileInfo struct {
    FilePath  string // file location
    Content   []byte // file content
    Sum       string // md5 sum of content
    IsRenamed bool   // indicator whether the particular file is renamed already or not
}

func main() {
    log.Println("start")
    start := time.Now()

    // pipeline 1: loop all files and read it
    chanFileContent := readFiles()
    fmt.Println("CHAN FILE CONTENT :",chanFileContent)
    fmt.Println("------------------------------------")

    // ...
	// pipeline 2: calculate md5sum
    chanFileSum1 := getSum(chanFileContent)
    chanFileSum2 := getSum(chanFileContent)
    chanFileSum3 := getSum(chanFileContent)
    fmt.Println("------------------------------------")
    chanFileSum := mergeChanFileInfo(chanFileSum1, chanFileSum2, chanFileSum3)
    fmt.Println("------------------------------------")

    // ...
	// pipeline 3: rename files
    chanRename1 := rename(chanFileSum)
    chanRename2 := rename(chanFileSum)
    chanRename3 := rename(chanFileSum)
    chanRename4 := rename(chanFileSum)
    fmt.Println("------------------------------------")
    chanRename := mergeChanFileInfo(chanRename1, chanRename2, chanRename3, chanRename4)
    fmt.Println("------------------------------------")

    // ...
	// print output
    counterRenamed := 0
    counterTotal := 0
    for fileInfo := range chanRename {
        if fileInfo.IsRenamed {
            counterRenamed++
        }
        counterTotal++
    }
    fmt.Println("------------------------------------")

    log.Printf("%d/%d files renamed", counterRenamed, counterTotal)

    duration := time.Since(start)
    log.Println("done in", duration.Seconds(), "seconds")
}

func readFiles() <-chan FileInfo {
    chanOut := make(chan FileInfo)

    go func() {
        fmt.Println("GO FUNC")
        err := filepath.Walk(tempPathThree, func(path string, info os.FileInfo, err error) error {
            fmt.Println(path)

            // if there is an error, return immediatelly
            if err != nil {
                return err
            }

            // if it is a sub directory, return immediatelly
            if info.IsDir() {
                return nil
            }

            buf, err := ioutil.ReadFile(path)
            if err != nil {
                return err
            }

            chanOut <- FileInfo{
                FilePath: path,
                Content:  buf,
            }

            return nil
        })
        if err != nil {
            log.Println("ERROR:", err.Error())
        }

        close(chanOut)
    }()
    fmt.Println("READ FILES")

    return chanOut
}

func getSum(chanIn <-chan FileInfo) <-chan FileInfo {
    chanOut := make(chan FileInfo)

    go func() {
        fmt.Println("CHAN IN :", chanIn)
        for fileInfo := range chanIn {
            fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))
            chanOut <- fileInfo
        }
        fmt.Println("CHAN OUT func getSum() :", chanOut)
        close(chanOut)
    }()

    return chanOut
}

func mergeChanFileInfo(chanInMany ...<-chan FileInfo) <-chan FileInfo {
    wg := new(sync.WaitGroup)
    chanOut := make(chan FileInfo)
    fmt.Println("CHAN IN MANY :", chanInMany)
    // fmt.Println("WG :", wg)

    wg.Add(len(chanInMany))
    for _, eachChan := range chanInMany {
        go func(eachChan <-chan FileInfo) {
            for eachChanData := range eachChan {
                chanOut <- eachChanData
            }
            fmt.Println("CHAN OUT func mergeChanFileInfo() :", chanOut)
            wg.Done()
        }(eachChan)
    }

    go func() {
        wg.Wait()
        close(chanOut)
    }()

    return chanOut
}

func rename(chanIn <-chan FileInfo) <-chan FileInfo {
    chanOut := make(chan FileInfo)

    go func() {
        for fileInfo := range chanIn {
            newPath := filepath.Join(tempPathThree, fmt.Sprintf("file-%s.txt", fileInfo.Sum))
            err := os.Rename(fileInfo.FilePath, newPath)
            fileInfo.IsRenamed = err == nil
            chanOut <- fileInfo
        }

        close(chanOut)
    }()

    return chanOut
}