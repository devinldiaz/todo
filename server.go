package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"


	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type todo struct {
	Item string
}

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
		rows.Scan(&res)            // use scan() to assign current row to res
		todos = append(todos, res) // append value of res to array
	}

	return c.Render("index", fiber.Map{
		"Todos": todos, // pass todos array to index view
	})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTodo := todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT into todos VALUES ($1)", newTodo.Item) // execute SQL query where we add new to-do item into DB
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}
 
	return c.Redirect("/") // redirect to home page
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	olditem := c.Query("olditem")
	newitem := c.Query("newitem")
	// replace old name with new one
	db.Exec("UPDATE todos SET item = $1 WHERE item = $2", newitem, olditem)
	return c.Redirect("/")
}
func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	todoToDelete := c.Query("item")
	db.Exec("DELETE from todos WHERE item=$1", todoToDelete)
	return c.SendString("Deleted")
}

func main() {

	connStr := "host=localhost user=postgres password=246Trin!trotoluene dbname=todo sslmode=disable"

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}


	/* configure our Fiber app to serve our HTML views */
	engine := html.New("./views", ".html")
	

	// create new Fiber object & assign to app variable
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

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
