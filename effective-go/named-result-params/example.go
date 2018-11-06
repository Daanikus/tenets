package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}

func (n *Node) Parent1() (node *Node)
func (n *Node) Parent2() (node *Node, err error)
