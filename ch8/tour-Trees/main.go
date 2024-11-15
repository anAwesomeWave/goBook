package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var recWalk func(curT *tree.Tree)
	recWalk = func(curT *tree.Tree) {
		if curT == nil {
			return
		}
		recWalk(curT.Left)
		ch <- curT.Value
		recWalk(curT.Right)
	}
	recWalk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for val1 := range ch1 {
		val2, ok := <-ch2
		if !ok {
			return false
		}
		if val1 != val2 {
			return false
		}
	}
	if _, ok := <-ch2; ok {
		return false
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
