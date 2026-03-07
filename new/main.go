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

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var products = []Product{
	{ID: "1", Name: "Laptop"},
	{ID: "2", Name: "Phone"},
}

func main() {

	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/products", getProducts)
		api.GET("/products/:id", getProduct)
		api.POST("/products", createProduct)
	}

	router.Run(":8080")
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {

	id := c.Param("id")

	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Product not found",
	})
}

func createProduct(c *gin.Context) {
	var newProduct Product

	c.BindJSON(&newProduct)

	products = append(products, newProduct)

	c.JSON(http.StatusCreated, newProduct)
}
