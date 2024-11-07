package main

import (
	"fmt"
	"goBook/ch7-Interfaces/1-word_lines_count/lib/counters"
)

func main() {
	var wc counters.WordCnt
	var lc counters.LinesCnt
	fmt.Fprintf(&wc, "my word 123")
	fmt.Fprintf(&wc, "aboba")
	fmt.Println(&wc)
	wc.Reset()
	fmt.Println(&wc)

	fmt.Fprintf(&lc, "my\nword 123")
	fmt.Fprintf(&lc, "aboba")
	fmt.Println(&lc)
	lc.Reset()
	fmt.Println(&lc)
}
