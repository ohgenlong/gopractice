package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("open file failed: %s", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func ReadBufIO(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("open file failed: %s", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file err: %s", err)
		}

		fmt.Printf("bufio read %s", line)
	}
}

func ReadSimple(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("open file failed: %s", err)
		return
	}
	defer file.Close()
	var data [10]byte
	for {
		n, err := file.Read(data[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file err: %s", err)
		}

		str := string(data[0:n])
		fmt.Printf("%s\n", str)
	}
}

func main() {
	filename := "D:/go/src/test.html"
	//ReadSimple(filename)
	//ReadBufIO(filename)
	ReadFile(filename)
}
