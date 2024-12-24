package main
import (
	"context"
	"fmt"
    "log"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func main(){
	app := fiber.New()
	app.Get("/", func(c*fiber.Ctx)error{
		return c.SendString("chaal gyo bhaya..")
	})
	app.Get("/hello",func(c*fiber.Ctx)error{
		return c.SendString("NAMASKARAM KYA HELLO ?")
	})
	app.Get("/hello/:id", func(c*fiber.Ctx)error{
		id := c.Params("id")
		return c.SendString("UserId:"+id)
	})

	type User struct{
		Name string `Json: "name"`
		Age int `Json: "age"`
	}
	app.Post("/user", func(c*fiber.Ctx)error{
		user := new(User)
		if err:= c.BodyParser(user); err != nil{
			return err
		}
		return c.JSON(user)
	})

	//settingup mongodb uri
	uri := "mongodb://localhost:27017/"

	//creating a new client and connect to server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
	}
	defer func(){
		if err = client.Disconnect(context.TODO()); err != nil{
			fmt.Println(err)
		}
	}()
	   // Ping the primary to verify connection
	   if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
        log.Fatal(err)
    }

	app.Listen(":3000")
}