//connStr := "host=localhost user=postgres password=246Trin!trotoluene dbname=todo sslmode=disable"

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// handlers accept fiber context object and pointer to DB connection
func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string
	rows, err := db.Query("SELECT * FROM todos") // use db object to execute SQL query - returns rows that match queries + errors

	defer rows.Close() // close rows

	if err != nil {
		log.Fatal(err)
		c.JSON("An error occurred")
	}
	for rows.Next() { // iterate over rows 
		rows.Scan(&res) // use scan() to assign current row to res
		todos = append(todos, res) // append value of res to array
	}

	return c.Render("index", fiber.Map{
		"Todos": todos, // pass todos array to index view
	})
}
func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}
func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}
func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func main() {
	connStr := "host=localhost user=postgres password=246Trin!trotoluene dbname=todo sslmode=disable"

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	/* configure our Fiber app to serve our HTML views */
	engine := html.New("./views",".html")
	// create new Fiber object & assign to app variable
	app := fiber.New(fiber.Config{
		Views: engine,
		}) 

	// ROUTES
	// + four handler methods that are called whenever someone visits those routes

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db) //pass the database connection to our handlers so we can use it to execute database queries
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT") // check environment variables for PORT
	if port == "" {
		port = "3000" // if port doesn't exist, set to 3000
	}

	// call app.listen to start HTTP server listening on port
	// log.Fatalln() to log the output to the console of any errors.
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
