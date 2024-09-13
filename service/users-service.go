package service

import (
	"acme/db"
	"acme/model"
	"errors"
	"fmt"
)

func GetUsers() ([]model.User, error) {
	users, err := db.GetUsers()

	if err != nil {
		fmt.Println("error getting users from db:", err)
		return nil, errors.New("there was an error getting the users from the database")
	}

	return users, nil
}

func CreateUser(newUser model.User) (int, error) {
	id, err := db.AddUser(newUser)

	if err != nil {
		fmt.Println("error adding new user to db:", err)
		return -1, errors.New("there was an error adding new user to the database")
	}

	return id, nil
}

func GetSingleUser(id int) (model.User, error) {
	fetchedUser, err := db.GetUser(id)
	if err != nil {
		fmt.Println("error fetching user from db:", err)
		return model.User{}, errors.New("error fetching user")
	}

	return fetchedUser, nil
}

func DeleteUser(id int) error {
	err := db.DeleteUser(id)

	if err != nil {
		fmt.Println("error deleting user:", err)
		return errors.New("error deleting user")
	}

	return nil
}

func UpdateUser(id int, body model.User) error {
	err := db.UpdateUser(id, body)

	if err != nil {
		fmt.Println("error updating user:", err)
		return errors.New("error updating user")
	}

	return nil
}
