package main

import "fmt"

type User struct {
	name  string
	score int
}

func main() {
	fmt.Println(User{name: "佐藤", score: 100})
}
