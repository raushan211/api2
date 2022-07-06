package main

import (
	"log"

	_ "github.com/lib/pq"
)

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

func createUserToDB(reqBody User) bool {
	var result = true
	sqlStatement := `
	INSERT INTO users(name, user_id, mobile, mail, city, password)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := DB.Exec(sqlStatement, reqBody.Name, reqBody.Mail, reqBody.Mobile, reqBody.UserID, reqBody.City, reqBody.Password)
	if err != nil {
		log.Fatal("error while inserting:", err)

	}

	return result
}

func updateUserFromDB(reqBody User, user_id string) bool {
	var result = true
	sqlStatement := `UPDATE users SET name = $2, mail=$3, city=$4, mobile=$5 WHERE id = $1`
	_, err := DB.Exec(sqlStatement, reqBody.ID, reqBody.Name, reqBody.Mail, reqBody.City, reqBody.Mobile)

	if err != nil {
		log.Fatal("Error in update: ", err)
	}

	return result
}
func deleteUserFromDB(userid string) bool {
	var result = true
	sqlStatement := `DELETE FROM users WHERE id = $1`
	_, err := DB.Exec(sqlStatement, userid)

	if err != nil {
		log.Fatal("Error in delete: ", err)
	}

	return result
}
