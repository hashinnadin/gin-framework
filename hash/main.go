// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/crypto/bcrypt"
// )

// var storedHash []byte

// func main() {

// 	router := gin.Default()

// 	router.POST("/register", registerHandler)
// 	router.POST("/login", loginHandler)
// 	router.GET("/logout", AuthMiddleware(),logoutHandler)

// 	router.GET("/dashboard", AuthMiddleware(),dashboardHandler)
// 	router.Run(":8080")
// }

// func registerHandler(c *gin.Context) {

// 	var data struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(400, gin.H{"error": "invalid json"})
// 		return
// 	}

// 	// input validation
// 	if data.Username == "" || data.Password == "" {
// 		c.JSON(400, gin.H{"error": "username and password required"})
// 		return
// 	}

// 	// hash password
// 	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "hashing failed"})
// 		return
// 	}

// 	storedHash = hash

// 	c.JSON(200, gin.H{
// 		"message": "user registered",
// 	})
// }

// func loginHandler(c *gin.Context) {

// 	var data struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(400, gin.H{"error": "invalid json"})
// 		return
// 	}

// 	// validate input
// 	if data.Username == "" || data.Password == "" {
// 		c.JSON(400, gin.H{"error": "username and password required"})
// 		return
// 	}

// 	// compare password with hash
// 	err := bcrypt.CompareHashAndPassword(storedHash, []byte(data.Password))
// 	if err != nil {
// 		c.JSON(401, gin.H{"error": "invalid password"})
// 		return
// 	}

// 	// create session cookie
// 	c.SetCookie("session", "loggedin", 3600, "/", "localhost", false, true)

// 	c.JSON(200, gin.H{
// 		"message": "login successful",
// 	})
// }
// func logoutHandler(c *gin.Context) {

// 	c.SetCookie("session", "", -1, "/", "localhost", false, true)

// 	c.JSON(200, gin.H{
// 		"message": "logged out",
// 	})
// }
// func dashboardHandler(c *gin.Context) {

// 	c.JSON(200, gin.H{
// 		"message": "welcome to dashboard",
// 	})
// }

// func AuthMiddleware() gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		session, err := c.Cookie("session")

// 		if err != nil || session != "loggedin" {
// 			c.JSON(401, gin.H{"error": "login required"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// // }

// package main

// import "github.com/gin-gonic/gin"

// type User struct {
// 	Username string `json:"name"`
// 	Age      int    `json:"age"`
// }

// func main() {
// 	router := gin.Default()

// 	router.GET("/hello", hello)
// 	router.POST("hello", postHello)
// 	router.Run(":8080")
// }

// func postHello(c *gin.Context) {
// 	var user User
// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "invalid json",
// 		})
// 		return
// 	}
// 	c.JSON(200, gin.H{
// 		"name": user.Username,
// 		"age":  user.Age,
// 	})
// }
// func hello(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "hello hashin",
// 	})
// }

package main

import "github.com/gin-gonic/gin"

type Data struct {
	Data string `json:"message"`
}

func main() {

	r := gin.Default()
	r.GET("/", hello)
	r.POST("/", helloPost)
	r.Run(":8080")
}

func helloPost(c *gin.Context) {
	var mes Data
	c.BindJSON(&mes)
	c.JSON(200, gin.H{
		"message": mes.Data,
	})
}
func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello hashin",
	})
}
