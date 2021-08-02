package validation

import (
	"strconv"

	"github.com/YogeshTembe/go_json/model"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

var UserIds = make(map[string]model.User)

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func Validate(user *model.User) bool {
	isValid := true
	if user.Name == "" {
		log.Warn(New("Name is empty for userId- " + user.Id.String()))
		isValid = false
	}
	if user.Email == "" {
		log.Warn(New("Email is empty for userId- " + user.Id.String()))
		isValid = false
	}
	if len(strconv.Itoa(user.PhoneNumber)) != 10 {
		log.Warn(New("Phone number is less than 10 digits for userId-" + user.Id.String()))
		isValid = false
	}
	if user.Id.String() == "00000000-0000-0000-0000-000000000000" {
		user.Id = uuid.NewV4()
	}
	_, isFound := UserIds[user.Id.String()]
	if isFound {
		newId := uuid.NewV4()
		log.Warn(New("userId- " + user.Id.String() + " for the user- " + user.Name + " is already present so we will update the id to-" + newId.String()))
		user.Id = newId
	}

	return (isValid)
}
