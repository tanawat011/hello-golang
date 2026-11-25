package main

import "github.com/tanawat011/hello-golang/school"

func main() {
	println("School Name", school.SchoolName)
	println("School Address:", school.GetSchoolAddress())
}
