package main

import (
	"github.com/YogeshTembe/go_json/file_handling"
)

func main() {
	csvFile, _ := file_handling.OpenCSVFile("emp.csv")
	defer file_handling.CloseCSVFile(csvFile)
	users := file_handling.ReadCSVFile(csvFile)
	//fmt.Println(users)
	file_handling.WriteJSONFile("emp.json", users)
}
