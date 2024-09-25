package main

import (
	"fmt"

	"goBook/ch2/myLib"
)


func main() {
	var start myTime.Time = 1.
	fmt.Println(myTime.Delta(myTime.Begin, start))
	fmt.Println(myTime.Delta(start, start))
}
