package main

import "fmt"

func main() {
	s := "hahsin"
	fmt.Println(s)
	b := []byte(s)
	b[0] = 'H'
	s = string(b)
	fmt.Println(s)
}
