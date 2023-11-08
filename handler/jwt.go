package handler

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitJWT() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}
	return []byte(os.Getenv("SECRET_KEY"))
}
