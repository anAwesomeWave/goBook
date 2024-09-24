package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "пропуск новой строки")
var sep = flag.String("s", " ", "separator")

func main() {

	flag.Parse()
	arr := []string{"123", "asdb", "qwerty"}

	fmt.Printf("DELIM: %v\n", *sep)
	fmt.Printf("N: %v\n", *n)

	fmt.Print(strings.Join(arr, *sep))

	if *n {
		fmt.Println()
	}
}
