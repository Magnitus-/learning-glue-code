package main

import (
	"fmt"
	"io"
	"os"
)

func getFilePath() string {
	if len(os.Args) < 2 {
		fmt.Println("Missing file argument")
		os.Exit(1)
	}
	return os.Args[1]
}

func getFileReader(path string) *os.File {
	fPtr, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return fPtr
}

func main() {
	fPath := getFilePath()
	fPtr := getFileReader(fPath)
	io.Copy(os.Stdout, fPtr)
}
