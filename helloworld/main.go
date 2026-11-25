package main

import "fmt"

var name = "Tanawat Pintongpan"
var address string

// fmt.Println writes to standard output; the builtin println used in 2018
// writes to standard error and is meant for debugging the runtime.
func main() {
	for i := 1; i < 9; i++ {
		fmt.Println(i)
	}
	fmt.Println("Hello World")
	printAddress()
	printMyName(name)

	address = getAddress()
	fmt.Println(address)

	resultAddress, resultCode := getAddressAndYear()
	fmt.Println(resultAddress)
	fmt.Println(resultCode)
}

func printAddress() {
	fmt.Println("กรุงเทพ")
}

func printMyName(myName string) {
	fmt.Println(myName)
}

func getAddress() string {
	return "นครศรีธรรมราช"
}

// getAddressAndYear returns two values at once.
func getAddressAndYear() (string, int) {
	address := "นครศรีธรรมราช"
	code := 2016
	return address, code
}
