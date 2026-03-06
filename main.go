// package main

// import "fmt"

//	func main() {
//		s := "hahsin"
//		fmt.Println(s)
//		b := []byte(s)
//		b[0] = 'H'
//		s = string(b)
//		fmt.Println(s)
//	}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "Hello, World!",
	}

	json.NewEncoder(w).Encode(response)
}

// POST /api/user
func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("User received:", user.Name, user.Age)

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": "User received successfully",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/user", userHandler)

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}
