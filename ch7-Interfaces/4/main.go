package main

import (
	"fmt"
	"io"
)

type BasicReader string

func (br *BasicReader) Read(p []byte) (int, error) {
	sz := len(p)
	var err error
	if sz > len(*br) {
		sz = len(*br)
		err = io.EOF
	}
	return copy(p, (*br)[:sz]), err
}

func main() {
	reader := BasicReader("my string")
	bytes, _ := io.ReadAll(&reader)
	fmt.Println(string(bytes))
}
