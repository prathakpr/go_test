package main
import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("chaal gyo bhaya..")
	})
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("NAMASKARAM KYA HELLO ?")
	})
	app.Get("/hello/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("UserId:" + id)
	})

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	app.Post("/user", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).SendString("Invalid request payload")
		}
		return c.JSON(user)
	})

	// MongoDB URI
	uri := "mongodb://localhost:27017/"

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Failed to disconnect:", err)
		}
	}()

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	// Access DB and collection
	collect := client.Database("test_go").Collection("human")

	// Insert a document
	doc := bson.D{{"name", "Alice"}, {"age", 30}}
	result, err := collect.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	app.Listen(":3000")
}
