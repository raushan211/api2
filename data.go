package main

func getUserByID(userID string) User {
	user := User{}
	if user, ok := Data[userID]; ok {
		return user
	}
	return user
}
