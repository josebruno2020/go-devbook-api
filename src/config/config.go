package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionDB = ""
	AppPort      = 0
)

func Setup() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		AppPort = 2000
	}

	AppPort = port
	ConnectionDB = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
