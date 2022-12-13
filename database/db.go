package database

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Balance  float64  `json:"balance"`
}

func getUsers() ([]User, error) {
	data, err := os.ReadFile("db.json")
	var users []User

	if err == nil {
		json.Unmarshal(data, &users)
	}

	return users, err
}

func updateDB(data []User) {
	bytes, err := json.Marshal(data)

	if err == nil {
		os.WriteFile("db.json", bytes, 0644)
	} else {
		panic(err)
	}
}

func FindUser(username string) (*User, error) {
	users, err := getUsers()
	if err == nil {
		for i := 0; i < len(users); i++ {
			user := users[i]
			if strings.EqualFold(user.Username, username) {
				return &user, nil
			}
		}
	}
	return nil, err
}

func FindOrCreateUser(username string) (*User, error) {
	user, err := FindUser(username)

	if user == nil {
		var newUser User
		newUser.Username = strings.ToLower(username)
		newUser.Balance = 0
		users, err := getUsers()
		if err == nil {
			users = append(users, newUser)
			updateDB(users)
		}
		fmt.Printf("Cannot find user %v \n", username)
		return &newUser, err
	}
	return user, err
}

func UpdateUser(user *User) {
	// Update the json with modified user info
	users, err := getUsers()
	if err == nil {
		for i := 0; i < len(users); i++ {
			if strings.EqualFold(users[i].Username, user.Username) {
				// Update users details
				users[i] = *user
			}
		}
		updateDB(users)
	}
}
