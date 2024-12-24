package main
import "github.com/gofiber/fiber/v2"

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
	app.Listen(":3000")
}