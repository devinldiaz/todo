package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2" // Import the Fiber framework
)

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func postHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func putHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func deleteHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func main() {
	app := fiber.New() // create new Fiber object & assign to app variable

	// add methods to handle GET, POST, PUT, and DELETE operations for app
	// + four handler methods that are called whenever someone visits those routes
	app.Get("/", indexHandler)

	app.Post("/", postHandler)

	app.Put("/update", putHandler)

	app.Delete("/delete", deleteHandler)

	port := os.Getenv("PORT") // check environment variables for PORT
	if port == "" {
		port = "3000" // if port doesn't exist, set to 3000
	}
	// call app.listen to start HTTP server listening on port
	// log.Fatalln() to log the output to the console of any errors.
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
