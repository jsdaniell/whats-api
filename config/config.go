package config

import (
	"log"
	"os"
	"strconv"
)

var (
	PORT = 9000
)

func Load() {
	var err error

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
		PORT = 9000
	}
}
