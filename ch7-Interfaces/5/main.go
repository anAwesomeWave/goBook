package main

import (
	"fmt"
	"io"
	"strings"
)

type LReader struct {
	r    io.Reader
	size int64
}

func (lr *LReader) Read(p []byte) (int, error) {
	var err error
	if int64(len(p)) > lr.size {
		err = io.EOF
		p = p[:lr.size]
		//p = make([]byte, lr.size)
		// p - копия, но указывает на тот участок памяти
		// куда посылаем. сделав срез выше, мы указываем туда же
		// создав новый слайс, начинаем указывать в другое место
	}
	n, rErr := lr.r.Read(p)
	if rErr != nil {
		err = rErr
	}
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LReader{r, n}
}
func main() {
	newReader := strings.NewReader("123456789")
	lReader := LimitReader(newReader, 2)
	for i := 0; i < 4; i++ {
		p := make([]byte, i)
		n, err := lReader.Read(p)
		fmt.Println(i, n, err, p)
	}
}
