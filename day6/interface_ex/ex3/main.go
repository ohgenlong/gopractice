package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	Name string
}

func (this *File) Read() {
	fmt.Println("read: ", this.Name)
}

func (this *File) Write() {
	fmt.Println("write: ", this.Name)
}

func Test(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {

	var f File
	f.Name = "testfile"
	Test(&f)

	var ff *File
	var b interface{}
	b = ff
	if sv, ok := b.(ReadWriter); ok {
		fmt.Printf("f implement ReadWriter: %s %b\n", sv, ok)
	}

}
