package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SHost string
	SPort string

	DBUser string
	DBPass string
	DBAdrs string
	DBName string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		SHost:  os.Getenv("SHost"),
		SPort:  os.Getenv("SPort"),
		DBUser: os.Getenv("DBUser"),
		DBPass: os.Getenv("DBPass"),
		DBAdrs: fmt.Sprintf("%s:%s", os.Getenv("DBHost"), os.Getenv("DBPort")),
		DBName: os.Getenv("DBName"),
	}
}
