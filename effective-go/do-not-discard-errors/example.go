package main

import (
    "fmt"
    "os"
)

func main() {
    var filename string = "path/to/file"
    file, err := os.Create(filename)
    file, _ = os.Create(filename)
    f, _ := os.Create(filename)
}