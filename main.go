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

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// // handler function
// func handler(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodGet {
// 		fmt.Println( "GET request received")

// 	} else if r.Method == http.MethodPost {
// 		fmt.Fprintln(w, "POST request received")

// 	} else {
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 	}
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		fmt.Fprintln(w, "Request receved")
// 	} else if r.Method == http.MethodPost {
// 		fmt.Fprintln(w, "succeefullu posted")
// 	} else {
// 		http.Error(w, "method not setteed", http.StatusMethodNotAllowed)
// 	}
// }

// func main() {

// 	http.HandleFunc("/api", handler)

// 	fmt.Println("Server running on port 8080")

// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"encoding/json"
	"net/http"
)

// handler for /api/hello
// func helloHandler(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	response := map[string]string{
// 		"message": "Hello, World",
// 	}

// 	json.NewEncoder(w).Encode(response)
// }

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Hello",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/api/hello", helloHandler)

	http.ListenAndServe(":8080", nil)
}
