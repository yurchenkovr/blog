package main

import (
	"fmt"
	"blog/api/postrges"
	"blog/router"
)

/*
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"-"`
}*/
func main() {
	fmt.Println("hello")
	postrges.New()
	//postrges.GetAllArticles(db)

	e := router.New()

	e.Logger.Fatal(e.Start(":8080"))
}
