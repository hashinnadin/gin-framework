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

// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	r := gin.Default()
// 	r.POST("/login", loginHandler)
// 	r.GET("/profile", profileHandler)
// 	r.GET("/logout", logoutHandler)
// 	r.Run(":8080")
// }

// func loginHandler(c *gin.Context) {
// 	var data struct {
// 		Username string `json:"username"`
// 		Password int    `json:"password"`
// 	}

// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
// 		return
// 	}
// 	if data.Username == "admin" && data.Password == 1234 {
// 		c.SetCookie("session", "loggedin", 3600, "/", "localhost", false, true)
// 		c.JSON(http.StatusOK, gin.H{"message": "Successfully logged"})
// 		return
// 	}
// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
// }

// func profileHandler(c *gin.Context) {
// 	session, err := c.Cookie("session")

// 	if err != nil || session != "loggedin" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "please login"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "welcome to homepage"})
// }

// func logoutHandler(c *gin.Context) {

// 	c.SetCookie("session", "", -1, "/", "localhost", false, true)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "logged out",
// 	})
// }

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		fmt.Println("Request:", c.Request.Method, c.Request.URL.Path)

		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func main() {

	router := gin.Default()

	router.POST("/login", LoggingMiddleware(), loginHandler)
	router.GET("/logout", LoggingMiddleware(), logoutHandler)

	router.Run(":8080")
}

func loginHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "login endpoint",
	})
}

func logoutHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "logout endpoint",
	})
}
