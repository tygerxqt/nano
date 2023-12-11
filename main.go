package main

import (
	"log"
	"nano/api"

	"github.com/fatih/color"
	"github.com/spf13/viper"
	bolt "go.etcd.io/bbolt"
)

func main() {
	info := color.New(color.Bold, color.FgBlue).PrintlnFunc()
	// warning := color.New(color.BgYellow).PrintfFunc()
	error := color.New(color.BgRed).PrintlnFunc()

	config()

	info("Nano v0.0.1 - Your personal cloud")

	db, err := bolt.Open(viper.GetString("database"), 0600, nil)

	if err != nil {
		error("Error opening database, %s", err)
	}

	defer db.Close()

	stopApi := api.Start(db)
	defer stopApi()
}

func config() {
	viper.SetDefault("port", "8080")
	viper.SetDefault("database", "nano.db")

	viper.SetConfigName("prod")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.Panic("Error reading config file: ", err)
	}
}
