package main

import (
	"fmt"
	"unsafe"
)

type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}
func main() {
	f := NewFile(10, "./struct.go")
	fmt.Println(f)
	size := unsafe.Sizeof(File{})
	fmt.Println(size)
}
