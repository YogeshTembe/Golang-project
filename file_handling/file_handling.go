package file_handling

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/YogeshTembe/go_json/model"
	"github.com/YogeshTembe/go_json/validation"
	uuid "github.com/satori/go.uuid"
)

func OpenCSVFile(fileDir string) (*os.File, error) {
	csvFile, err := os.Open(fileDir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	return csvFile, err
}

func CloseCSVFile(csvFile *os.File) {
	csvFile.Close()
}

func ReadCSVFile(csvFile *os.File) []model.User {
	var users []model.User

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		id, _ := uuid.FromString(line[0])
		phoneNo, _ := strconv.Atoi(line[3])
		isActive, _ := strconv.ParseBool(line[4])

		user := model.User{
			Id:          id,
			Name:        line[1],
			Email:       line[2],
			PhoneNumber: phoneNo,
			IsActive:    isActive,
		}
		isValid := validation.Validate(&user)

		if isValid {
			users = append(users, user)
			validation.UserIds[user.Id.String()] = user
		}
	}
	fmt.Println("CSV file reading and data validation is done.")
	return users
}

func WriteJSONFile(fileDir string, users []model.User) {
	file, _ := json.MarshalIndent(users, "", " ")
	_ = ioutil.WriteFile(fileDir, file, 0644)
	fmt.Println("JSON file writing is done.")
}
