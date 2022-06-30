package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserID   string `json:"user_id"`
	Mobile   string `json:"mobile"`
	Mail     string `json:"mail"`
	City     string `json:"city"`
	Password string `json:"password" binding:"required"`
}

var (
	Data map[string]User
	DB   *sql.DB
)

func main() {
	createDBConnection()
	defer DB.Close()
	Data = make(map[string]User)
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func setupRoutes(r *gin.Engine) {
	r.GET("/user/:user_id", GetUserById)
	r.GET("/user", GetAllUser)
	r.PUT("/user/:user_id", UpdateUser)
	r.POST("/user", CreateUser)
	r.DELETE("/user/:user_id", DeleteUser)
}
func GetUserById(c *gin.Context) {
	userID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{
			"error": "user id is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	user, _ := getUserByIDFromDB(userID)
	res := gin.H{
		"user": user,
	}
	c.JSON(http.StatusOK, res)
}
func GetAllUser(c *gin.Context) {
	users, err := getAllUserFromDB()
	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := gin.H{
		"users": users,
	}
	c.JSON(http.StatusOK, res)
}
func UpdateUser(c *gin.Context) {
	userID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{
			"error": "user id is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	reqBody := User{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if reqBody.UserID == "" {
		res := gin.H{
			"error": "UserID can't be empty",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if reqBody.UserID != userID {
		res := gin.H{
			"errror": "UserID can't be updated",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	password := c.GetHeader("password")
	userObj := getUserByID(userID)
	if userObj.UserID == "" {
		res := gin.H{
			"error": "UserID can't be empty"}
		c.JSON(http.StatusBadRequest, res)
		return

	}
	if password != Data[userID].Password {
		res := gin.H{
			"errror": "incorrect password",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Data[userID] = reqBody
	res := gin.H{
		"success": true,
		"user_id": reqBody,
	}
	c.JSON(http.StatusOK, res)
	return
}
func CreateUser(c *gin.Context) {
	reqBody := User{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if reqBody.UserID == "" {
		res := gin.H{
			"error": "UserId must not be empty",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if len(reqBody.Mobile) != 10 {
		res := gin.H{
			"error": "phone number must be 10 digit",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Data[reqBody.UserID] = reqBody
	res := gin.H{
		"success": true,
		"user":    reqBody,
	}
	c.JSON(http.StatusOK, res)
	return
}
func DeleteUser(c *gin.Context) {
	userID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{
			"error": "user_id is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	if _, ok := Data[userID]; ok {
		delete(Data, userID)

		res := gin.H{
			"success": true,
		}
		c.JSON(http.StatusOK, res)
		return
	}

	res := gin.H{
		"error": "user_id doesnot exist",
	}
	c.JSON(http.StatusBadRequest, res)
}
