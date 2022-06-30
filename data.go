package main

import (
	"log"

	_ "github.com/lib/pq"
)

func getUserByID(userID string) User {
	user := User{}
	if user, ok := Data[userID]; ok {
		return user
	}
	return user
}

func getUserByIDFromDB(userID string) (User, error) {

	// defer DB.Close()
	user := User{}
	userSQL := "SELECT id, name, user_id, mobile, city, mail, password FROM users WHERE user_id = $1"

	err := DB.QueryRow(userSQL, userID).Scan(&user.ID, &user.Name, &user.UserID, &user.Mobile, &user.City, &user.Mail, &user.Password)
	if err != nil {
		log.Println("Failed to execute query: ", err)
	}
	return user, err
}

func getAllUserFromDB() ([]User, error) {
	users := []User{}
	userSQL := "SELECT id, name, user_id, mobile, city, mail, password FROM users"
	rows, err := DB.Query(userSQL)
	if err != nil {
		log.Println("Failed to execute query: ", err)
		return users, err
	}
	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Name, &user.UserID, &user.Mobile, &user.City, &user.Mail, &user.Password)
		users = append(users, user)
	}
	return users, err
}
