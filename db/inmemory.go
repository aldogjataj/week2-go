package db

import (
	"acme/model"
	"errors"
	"slices"
)

var users []model.User
var count int = 3

func init() {
	// Initialize the in-memory database with some sample data
	users = []model.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}
}

func GetUsers() ([]model.User, error) {
	return users, nil
}

func AddUser(user model.User) (int, error) {
	count++
	user.ID = count
	users = append(users, user)
	return user.ID, nil
}

func GetUser(id int) (model.User, error) {
	var user model.User

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return user, errors.New("user id not found")
}

func DeleteUser(id int) error {
	for index, user := range users {
		if user.ID == id {
			users = slices.Delete(users, index, index+1)
			return nil
		}
	}

	return errors.New("could not delete user with no id found")
}

func UpdateUser(id int, body model.User) error {
	for index, user := range users {
		if user.ID == id {
			users[index].Name = body.Name
			// userPointer := &users[index]
			// userPointer.Name = body.Name
			return nil
		}
	}

	return errors.New("could not update user with no id found")
}
