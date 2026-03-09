// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Product struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

// var products = []Product{
// 	{ID: "1", Name: "Laptop"},
// 	{ID: "2", Name: "Bike"},
// }

// func main() {

// 	router := gin.Default()

// 	api := router.Group("/api")
// 	{
// 		api.GET("/products", getProducts)
// 		// api.GET("/products/:id")
// 		// api.POST("/products")
// 	}
// 	router.Run(":8080")
// }
// func getProducts(c *gin.Context) {
// 	c.JSON(http.StatusOK, products)
// }

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
// 	{ID: 2, Name: "Shahala"},
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

// 	err := c.ShouldBindJSON(&newUser)

// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"Error": "Invalid json",
// 		})
// 		return
// 	}
// 	users = append(users, newUser)
// 	c.JSON(201, newUser)
// }

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", loginHandler)
	router.Run(":8080")
}

var data struct {
	Username string `json:"username"`
	Password int    `json:"password"`
}

func loginHandler(c *gin.Context) {

	err := c.ShouldBindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}
	if data.Username == "admin" && data.Password == 1234 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully login",
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "invalid username or password",
	})
}
