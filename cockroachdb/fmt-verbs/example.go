package main

import (
	"fmt"
)

func main() {
	var s string = "Hello, world!"
	var i int = 42
	fmt.Println("%v", s)
	fmt.Println("%s", s)
	fmt.Println("%v%v", s, i)
	fmt.Println("%v", i)
	fmt.Println("%v is a string", s)
}
