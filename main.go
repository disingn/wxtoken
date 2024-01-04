package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"wxlogin/cfg"
	"wxlogin/dao"
	"wxlogin/sever"
)

func init() {
	cfg.Init()
	dao.Init()
}
func main() {
	app := fiber.New()
	app.Use(cors.New(cors.ConfigDefault))
	app.Use(logger.New(logger.ConfigDefault))
	app.Post("/wx", func(c *fiber.Ctx) error {
		sever.WXMsgReceive(c)
		return nil
	})
	app.Get("/wx", func(c *fiber.Ctx) error {
		sever.WXCheckSignature(c)
		return nil
	})
	err := app.Listen(":3100")
	if err != nil {
		return
	}

	//k, _ := sever.SetToken()
	//fmt.Printf("%s", k)

}
