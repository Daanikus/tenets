package main

import (
	"fmt"
)

func  main() {
	fmt.Println("Hello, playground")
	// This comment not on a func decl
}

/* foo is the first word
 */
func foo() {}

/* This func comment should begin with 'bar'
 */
func bar() {}

// This func comment should begin with 'baz'
// and we should not worry about this line
func baz() {}
