package handlers

import (
	"context"
	"github.com/pauljc-probeplus/todo_app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RenderLogin(c *fiber.Ctx) error {
	return c.Render("login", nil)
}

func HandleLogin(c *fiber.Ctx, db *mongo.Database) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).SendString("Invalid Input")
	}

	var user models.User
	filter := bson.M{"user_name": input.Username, "password": input.Password}
	err := db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return c.Status(401).SendString("Invalid Username or Password")
	}

	return c.Redirect("/home?user_name=" + input.Username)
}
