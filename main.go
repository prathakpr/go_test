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
	app.Listen(":3000")
}