package main

import (
	"fmt"
	"io"
	"os"
)

type MyCountingWriter struct {
	wr  io.Writer
	cnt int64
}

func (mcw *MyCountingWriter) Write(p []byte) (int, error) {
	n, err := mcw.wr.Write(p)
	mcw.cnt += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := MyCountingWriter{w, 0}
	return &newWriter, &newWriter.cnt
}

func main() {
	wr, cnt := CountingWriter(os.Stdout)
	fmt.Println("abc")       // default writer, non counting
	fmt.Fprintf(wr, "abc\n") // my counting writer
	fmt.Println(*cnt)

}
