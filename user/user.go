package user

import (
	"database/sql"
	"fmt"

	"github.com/andrew-lawlor/cc-simple-auth-server/db"
)

type User struct {
	UserID      int    `json:"userID"`
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password,omitempty"`
	Email       string `json:"email,omitempty"`
	Created     string `json:"created,omitempty"`
}

func NewUser(userName, displayName, hashedPassword, email string) bool {
	var user = User{
		UserName:    userName,
		DisplayName: displayName,
		Password:    hashedPassword,
		Email:       email,
	}
	return createUser(user)
}

func createUser(user User) bool {
	db := db.GetDB()
	// Create
	statement, err := db.Prepare("INSERT INTO users (userName, password, displayName, email) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	_, err = statement.Exec(user.UserName, user.Password, user.DisplayName, user.Email)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func GetUser(userName string) (User, error) {
	db := db.GetDB()
	stmt, err := db.Prepare("SELECT userID, userName, password, displayName, created FROM users WHERE userName = ?")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	var user User
	err = stmt.QueryRow(userName).Scan(&user.UserID, &user.UserName, &user.Password, &user.DisplayName, &user.Created)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case of no rows returned.
			return user, err
		}
	}
	return user, err
}
