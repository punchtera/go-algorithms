package main

import (
	"filepath"
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	filename := filepath.Base(os.Args[0])
	fmt.Printf("filename %s \n", filename)

	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Printf("[%d] : [%s]\n", i, arg)
	}
	fmt.Println(s)
}
