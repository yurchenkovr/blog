package main

import (
	"blog/src/infrastructure/transport"
	"fmt"
)

func main() {
	fmt.Println("hello")
	e := transport.New()

	e.Logger.Fatal(e.Start(":8080"))
}
