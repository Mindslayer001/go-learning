package main

import (
	"fmt"
	"os"

	"github.com/mindslayer001/go-project-1/simpleMaths"
)

func main() {
	fmt.Println("The main is running")
	arguments := os.Args[1:]
	simpleMaths.Exparser(arguments[0])
}
