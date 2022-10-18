// CRUD
// CREATE
// READ
// UPDATE
// DELETE

// MVC
// Model View Controller

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MODELS
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Elo  string `json:"elo"`
}

type ResponseModel struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Err     bool   `json:"error"`
	Message string `json:"message"`
}

// DATABASE
var users = []User{
	{ID: "1", Name: "Axel", Age: 22, Elo: "Platinum II"},
	{ID: "2", Name: "Fabian", Age: 23, Elo: "Platinum IV"},
	{ID: "3", Name: "Nicolas", Age: 23, Elo: "Platinum III"},
}

// SERVER
func main() {
  router := gin.Default()

  // ROUTES
  router.GET("/api/getUsers", getUsers)
  router.GET("/api/getUser/:id", getUser)
  router.POST("/api/createUser", createUser)
  router.PUT("/api/setUser/:id", setUser)
  router.DELETE("/api/deleteUser/:id", deleteUser)

  router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// CONTROLLERS
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}

	response := ResponseModel{Status: true, Code: 404, Err: true, Message: "User not found"}

	c.IndentedJSON(http.StatusNotFound, response)
}

func createUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// VALIDATIONS & MIDDLEWARES
	for _, user := range users {
		if user.ID == newUser.ID {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "User already exists by ID",
			})
			return
		}
	}

	users = append(users, newUser)
	
	c.IndentedJSON(http.StatusCreated, newUser)
}

func setUser(c *gin.Context) {
	id := c.Param("id")

	var updatedUser User

	if err := c.BindJSON(&updatedUser); err != nil {
		return
	}

	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			c.IndentedJSON(http.StatusOK, users[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}

func RemoveAtIndex(s []User, index int) []User {
	return append(s[: index], s[index + 1 :]...)
}

func deleteUser(c *gin.Context) {
   id := c.Param("id")

	for i, user := range users {
		if user.ID == id {
			users = RemoveAtIndex(users, i)
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "User successfully deleted",
			})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}