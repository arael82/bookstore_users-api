package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		//TODO: Handle properly
		fmt.Println("Error:", err)
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO: Handle properly
		fmt.Println("Error:", err)
		return
	}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		//TODO: Handle properly
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
