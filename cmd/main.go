package main

import (
	"log"

	"github.com/emersonjsouza/go-api/pkg/common/config"
	"github.com/emersonjsouza/go-api/pkg/common/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()
	db.Init(c.DBUrl)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(c.Port)
	})

	app.Listen(c.Port)
}
