package Pathparser

import (
	"fmt"
	"os"
)

func Parser(fp string) string {
	if fp == "" {
		return "the file path is not mentioned"
	}

	fmt.Printf("the file path is %s", fp)
	var filepath []string
	var err error
	filepath, err = os.ReadDir(fp)
	if err != nil {
		fmt.Printf("the file path has error that is %s", err)
	}
	fmt.Printf("the file path retrived by is %s", filepath)

	return ""
}
