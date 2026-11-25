package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	run(os.Stdout)
}

// run shows that a deferred call waits until the surrounding function
// returns: "finish" is printed last although it is deferred before "second".
func run(w io.Writer) {
	printFirst(w)
	defer printFinish(w)
	printSecond(w)
}

func printFirst(w io.Writer) {
	fmt.Fprintln(w, "first")
}

func printSecond(w io.Writer) {
	fmt.Fprintln(w, "second")
}

func printFinish(w io.Writer) {
	fmt.Fprintln(w, "finish")
}
