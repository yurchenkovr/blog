package main

import (
	"blog/src/infrastructure/router"
	"fmt"
)

func main() {
	fmt.Println("hello")
	e := router.New()

	e.Logger.Fatal(e.Start(":8080"))
}
