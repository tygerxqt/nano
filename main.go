package main

import (
	"nano/api"
	"nano/logger"
	"nano/ui"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/spf13/viper"
	"go.etcd.io/bbolt"
)

func main() {
	config()
	initalise()

	db, err := bbolt.Open("./"+viper.GetString("data_dir")+"/"+viper.GetString("database"), 0600, nil)
	if err != nil {
		logger.Error("Error opening database: ", err)
	}

	defer db.Close()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.All("/*", filesystem.New(filesystem.Config{
		Root: ui.Dist(),
	}))

	api.CreateApiRoutes(app, db)

	logger.Subtle("Running on http://localhost:" + viper.GetString("port"))

	defer app.Shutdown()
	app.Listen(":" + viper.GetString("port"))
}

func config() {
	viper.SetDefault("port", "8080")
	viper.SetDefault("database", "nano.db")
	viper.SetDefault("data_dir", "nano-data")

	viper.SetConfigName("prod")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.Panic("Error reading config file: ", err)
	}
}

func initalise() {
	if _, err := os.Stat("./" + viper.GetString("data_dir")); os.IsNotExist(err) {
		err := os.Mkdir("./"+viper.GetString("data_dir"), 0755)

		if err != nil {
			logger.Error("Error creating nano-data directory: ", err)
		}
	}

	logger.Info("Nano v0.0.1 - Your personal cloud")
}
