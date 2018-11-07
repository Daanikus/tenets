package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create(filename)
    file, _ = os.Create(filename)
    f, _ := os.Create(filename)
}