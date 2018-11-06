package main

import (
	"fmt"
)

type person struct {
    name string
    age  int
}

func main() {
	fmt.Println("Hello, playground")
  //person{"Bob", 20}
  person{age: 20}
}
