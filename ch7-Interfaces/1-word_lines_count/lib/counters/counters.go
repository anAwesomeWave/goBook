package counters

import (
	"bufio"
	"fmt"
)

type baseCounter struct {
	x    int
	name string
}

func (w *baseCounter) Reset() {
	w.x = 0
}

func (w *baseCounter) String() string {
	return fmt.Sprintf("Total %s count: %d", w.name, w.x)
}

type WordCnt struct {
	baseCounter
}

func (w *WordCnt) Write(p []byte) (int, error) {
	w.name = "words" // hack
	isStop := false
	var err error
	for !isStop {
		nextBytePos, word, err := bufio.ScanWords(p, true) // eof == ' '
		p = p[nextBytePos:]
		isStop = len(word) == 0 || err != nil
		if !isStop {
			w.x += 1
		}
	}
	return len(p), err
}

type LinesCnt struct {
	baseCounter
}

func (w *LinesCnt) Write(p []byte) (int, error) {
	w.name = "lines" // hack
	isStop := false
	var err error
	for !isStop {
		nextBytePos, word, err := bufio.ScanLines(p, true) // eof == ' '
		p = p[nextBytePos:]
		isStop = len(word) == 0 || err != nil
		if !isStop {
			w.x += 1
		}
	}
	return len(p), err
}
