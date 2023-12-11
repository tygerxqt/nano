package api

import (
	"log"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.etcd.io/bbolt"
)

func Start(db *bbolt.DB) func() error {
	app := fiber.New(fiber.Config{
		AppName:               "Nano",
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	c := color.New(color.FgWhite, color.Faint, color.Italic).PrintlnFunc()
	c("[API] Listening on: http://localhost:", viper.GetString("port"))

	if err := app.Listen(":" + viper.GetString("port")); err != nil {
		log.Panic("Error starting API: ", err)
	}

	return app.Shutdown
}
