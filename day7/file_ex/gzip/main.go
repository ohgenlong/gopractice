package main

import (
	"os"
	"fmt"
	"compress/gzip"
	"bufio"
	"io"
)

func GzipTest(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("open file failed: %s", err)
		return
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Printf("gzip failed, err:%s\n", err)
	}
	bufreader := bufio.NewReader(reader)
	for {
		line, err := bufreader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf( "read err %s", err)
			return
		}
		fmt.Printf("%s", line)

	}

}

func main() {
	filename := "D:/go/src/test.gz"
	GzipTest(filename)
}
