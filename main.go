package main

import (
	"log"
	//"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/pauljc-probeplus/todo_app/config"
	"github.com/pauljc-probeplus/todo_app/routes"
)

func main() {
	// Initialize Fiber app with templates
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Initialize the database
	db := config.ConnectDB()
	defer config.DisconnectDB(db)

	// Register routes
	routes.RegisterRoutes(app, db)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
