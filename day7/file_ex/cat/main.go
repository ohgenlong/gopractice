package main

import (
	"os"
	"fmt"
	"io"
)

func Cat(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open file failed: %s\n", err)
		return
	}

	defer file.Close()
	io.Copy(os.Stdout, file)
}


func main() {

	if len(os.Args) == 1 {
		fmt.Printf("filename is requried")
		return
	}

	for i:=1; i<len(os.Args); i++ {
		Cat(os.Args[i])
	}
}
