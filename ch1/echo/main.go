package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func badArgsConcat(args []string) int{
	var sep, ans string
	for _, arg := range args {
		ans += sep + arg
		sep = " "
	}
	return len(ans)
}

func fastArgsConcat(args []string) int {
	return len(strings.Join(args, " "))
}

func main() {
	// fmt.Println(os.Args)
	fmt.Printf("Test on %d args\n", len(os.Args))

	fmt.Println("Bad concat function")
	timer := time.Now()

	fmt.Println(badArgsConcat(os.Args))

	fmt.Printf("Time elapsed: %d\n", time.Since(timer).Nanoseconds())

	fmt.Println("Fast concat function")
	timer = time.Now()

	fmt.Println(fastArgsConcat(os.Args))

	fmt.Printf("Time elapsed: %d\n", time.Since(timer).Nanoseconds())

}
