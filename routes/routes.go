package routes

import (
	"github.com/pauljc-probeplus/todo_app/handlers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(app *fiber.App, db *mongo.Client) {
	database := db.Database("todo_db")

	app.Get("/", handlers.RenderLogin)
	app.Post("/login", func(c *fiber.Ctx) error { return handlers.HandleLogin(c, database) })
	app.Get("/home", func(c *fiber.Ctx) error { return handlers.RenderHome(c, database) })
	app.Post("/create-task", func(c *fiber.Ctx) error { return handlers.CreateTask(c, database) })
	//app.Post("/update-task", func(c *fiber.Ctx) error { return handlers.UpdateTask(c, database) })
	app.Get("/edit-task", func(c *fiber.Ctx) error {return handlers.RenderEditTask(c, database)})
	app.Post("/save-task", func(c *fiber.Ctx) error {return handlers.SaveTask(c, database)})
	
}
