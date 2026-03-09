// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// var users = []User{
// 	{ID: 1, Name: "Hashin"},
// 	{ID: 2, Name: "Rahul"},
// }

// func main() {

// 	router := gin.Default()

// 	api := router.Group("/api")
// 	{
// 		api.GET("/users", getUsers)
// 		api.POST("/users", createUser)
// 	}

// 	router.Run(":8080")
// }

// func getUsers(c *gin.Context) {
// 	c.JSON(http.StatusOK, users)
// }

// func createUser(c *gin.Context) {

// 	var newUser User

// 	c.BindJSON(&newUser)

// 	users = append(users, newUser)

// 	c.JSON(http.StatusCreated, newUser)
// }

package main

func main() {

}
