package main

import (
	"fmt"
)

func main() {
	printFirst()
	defer printFinish()
	printSecond()
}

func printFirst() {
	fmt.Println("first")
}

func printSecond() {
	fmt.Println("second")
}
func printFinish() {
	fmt.Println("finish")
}
