package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"min=8"`
	Name     string `json:"name" validate:"required"`
	Status   bool   `json:"-"`
}

type Users []*User

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

var userList = Users{
	{
		ID:       1,
		Username: "dummy",
		Password: "dummy",
		Name:     "dummy",
		Status:   true,
	}, {
		ID:       2,
		Username: "caleb",
		Password: "caleb",
		Name:     "Caleb Manzano",
		Status:   true,
	},
}

func getNextID() int {
	return userList[len(userList)-1].ID + 1
}

var ErrUserNotFound = fmt.Errorf("User not found")

func findUserByID(id int) (*User, int, error) {
	for i, u := range userList {
		if id == u.ID {
			return u, i, nil
		}
	}

	return nil, -1, ErrUserNotFound
}

func GetUsers() Users {
	return userList
}

func AddUser(u *User) {
	u.ID = getNextID()
	userList = append(userList, u)
}

func UpdateUser(id int, u *User) error {
	_, pos, err := findUserByID(id)
	if err != nil {
		return err
	}

	u.ID = id
	userList[pos] = u
	return nil
}
