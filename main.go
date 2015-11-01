package main

import (
	"bufio"
	"os"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fi, err := os.Open("test_src.asm")
	Check(err)
	fr := bufio.NewReader(fi)
	Parse(fr)
	// next?
}
