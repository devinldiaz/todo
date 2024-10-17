package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2" // Import the Fiber framework
)

func main() {
	app := fiber.New() // create new Fiber object & assign to App variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
