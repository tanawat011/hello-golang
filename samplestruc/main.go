package main

import "fmt"
import "encoding/json" // มีการ Import Encoding JSON เข้ามา
type customer struct { // มีการปรับ Field ให้ขึ้นต้นตัวใหญ่
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func main() {
	customerLists := []customer{}
	customer1 := customer{
		Firstname: "Tanawat",
		Lastname:  "Pintongpan",
		Code:      111111,
		Phone:     "0000000000",
	}
	customer2 := customer{
		Firstname: "Khanungnit",
		Lastname:  "Pintongpan",
		Code:      222222,
		Phone:     "1111111111",
	}
	customerLists = append(customerLists, customer1)
	customerLists = append(customerLists, customer2)
	customerListsJson, err := json.Marshal(customerLists)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(customerListsJson))
}
