package handlers

import (
	"context"
	"github.com/pauljc-probeplus/todo_app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	//"bytes"
	//"net/http"
)

func RenderHome(c *fiber.Ctx, db *mongo.Database) error {
	username := c.Query("user_name")
	filter := bson.M{"user_name": username}

	log.Printf("Querying tasks with filter: %+v\n", filter)

	
	cursor, err := db.Collection("tasks").Find(context.Background(), filter)
	if err != nil {
		return c.Status(500).SendString("Error fetching tasks")
	}

	var tasks []models.Task
	if err := cursor.All(context.Background(), &tasks); err != nil {
		return c.Status(500).SendString("Error parsing tasks")
	}

	data := fiber.Map{
		"Username": username,
		"Tasks":    tasks,
	}
	return c.Render("home", data)
}

func CreateTask(c *fiber.Ctx, db *mongo.Database) error {
	
	log.Printf("Raw Request Body: %s\n", string(c.Body()))

	
	
	
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).SendString("Invalid Input")
	}
	log.Printf("Parsed Task: %+v\n", task)
	_, err := db.Collection("tasks").InsertOne(context.Background(), task)
	if err != nil {
		return c.Status(500).SendString("Error creating task")
	}

	log.Printf("Redirecting to: /home?user_name=%s\n", task.Username)

	return c.Redirect("/home?user_name=" + task.Username)
}

/*func UpdateTask(c *fiber.Ctx, db *mongo.Database) error {
	taskName := c.FormValue("task_name")
	username := c.FormValue("user_name")
	status := c.FormValue("status") == "on"

	filter := bson.M{"user_name": username, "task_name": taskName}
	update := bson.M{"$set": bson.M{"status": status}}

	_, err := db.Collection("tasks").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(500).SendString("Error updating task")
	}

	return c.Redirect("/home?user_name=" + username)
}*/




func RenderEditTask(c *fiber.Ctx, db *mongo.Database) error {
    username := c.Query("user_name")
    taskName := c.Query("task_name")

    // Fetch the task from the database
    var task models.Task
    filter := bson.M{"user_name": username, "task_name": taskName}
    err := db.Collection("tasks").FindOne(context.Background(), filter).Decode(&task)
    if err != nil {
        return c.Status(404).SendString("Task not found")
    }

    // Render the edit page with task details
    data := fiber.Map{
        "Username":  username,
        "TaskName":  task.TaskName,
        "Status":    task.Status,
    }
    return c.Render("edit-task", data)
}


func SaveTask(c *fiber.Ctx, db *mongo.Database) error {
    var task models.Task
    if err := c.BodyParser(&task); err != nil {
        return c.Status(400).SendString("Invalid input")
    }

    // Update the task in the database
    filter := bson.M{"user_name": task.Username, "task_name": c.FormValue("original_task_name")}
    update := bson.M{
        "$set": bson.M{
            "task_name": task.TaskName,
            "status":    task.Status,
        },
    }

    _, err := db.Collection("tasks").UpdateOne(context.Background(), filter, update)
    if err != nil {
        return c.Status(500).SendString("Failed to update task")
    }

    // Redirect back to the home page
    return c.Redirect("/home?user_name=" + task.Username)
}


