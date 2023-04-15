package main

import (
	"fmt"

	"github.com/arjun/modules/go-echo-api/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Welcome to ECHO API")

	// can be in router file/folder
	e := echo.New()

	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers.PostSingleHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
