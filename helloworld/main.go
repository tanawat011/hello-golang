package main

var name = "Tanawat Pintongpan"
var address string

func main() {
	for i := 1; i < 9; i++ {
		println(i)
	}
	println("Hello World")
	printAddress()
	printMyName(name)

	address = getAddress()
	println(address)

	resultCode, resultAddress := getAddressAndYear()
	println(resultCode)
	println(resultAddress)
}

func printAddress() {
	println("กรุงเทพ")
}

func printMyName(myName string) {
	println(myName)
}

func getAddress() string {
	return "นครศรีธรรมราช"
}

func getAddressAndYear() (string, int) {
	address := "นครศรีธรรมราช"
	code := 2016
	return address, code
}
