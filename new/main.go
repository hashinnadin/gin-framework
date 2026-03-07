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

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Hashin"},
	{ID: 2, Name: "Rahul"},
}

func main() {

	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/users", getUsers)
		api.POST("/users", createUser)
	}

	router.Run(":8080")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser User

	c.BindJSON(&newUser)

	users = append(users, newUser)

	c.JSON(http.StatusCreated, newUser)
}


arr:=int{1,2,0,8,4,56,0,90,74,0}
