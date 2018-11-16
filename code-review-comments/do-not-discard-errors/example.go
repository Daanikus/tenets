package main

import (
    "fmt"
    "os"
    "errors"
)

func main() {
    var filename string = "path/to/file"
    file, err := os.Create(filename)
    file, _ = os.Create(filename)
    f, _ := os.Create(filename)

    // Error deliberately discarded
    file, _ = os.Create(filename)

    // Error deliberately discarded
    var thing, _ = canThrow()
    fmt.Println("Done")
}

func canThrow() (int, error) {
    return 12, errors.New("some error")
}