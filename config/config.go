package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	PORT = 9000
)

func Load() {

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
		PORT = 9000
	}
}
