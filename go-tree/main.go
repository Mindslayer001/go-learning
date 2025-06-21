package main

import (
	"fmt"
	"os"

	Pathparser "github.com/mindslayer001/go-tree/PathParser"
)

func main() {
	fmt.Println("Hi")
	var Fp string = os.Args[1]
	var result string = Pathparser.Parser(Fp)
	fmt.Println(result)
}
