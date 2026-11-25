package main

import (
	"fmt"

	"github.com/tanawat011/hello-golang/school"
)

func main() {
	fmt.Println("School Name:", school.SchoolName)
	fmt.Println("School Address:", school.GetSchoolAddress())
}
