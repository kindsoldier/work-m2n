
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
  
)

func PathLength(path string) int {
    path = filepath.Clean(path)
    if len(path) == 0 {
        return 0
    }
    return len(strings.Split(path, "/"))
}


func PathWalkDir(basePath string, depth int) ([]string, error) {
    depth = depth + PathLength(basePath)
    var list []string
    err := filepath.Walk(basePath,
        func(filePath string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if PathLength(filePath) > depth  {
                return filepath.SkipDir
            }
            if info.IsDir() {
                list = append(list, filePath)
            }
            return nil
        })
    if err != nil {
        return list, err
    }
    return list, err
}


func main() {
    list, _ := PathWalkDir("/dat2", 5)
    fmt.Println(list)
}
