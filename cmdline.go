// +build ignore

package main

import (
	"fmt"
)

func main() {
	// p := NewProcessor()
	// fmt.Printf("%v\n", p)
	// p.LSR(0x0010)
	// fmt.Printf("%v\n", p)
	toks := Tokenize("LABEL CMP $0101 ;commenting stuff")
	toks, i := StripComments(toks)
	fmt.Printf("%v %d", toks, i)
}
