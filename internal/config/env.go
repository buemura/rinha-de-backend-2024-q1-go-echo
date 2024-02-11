package config

import "os"

var (
	PORT         string
	DATABASE_URL string
)

func LoadEnv() {
	PORT = os.Getenv("PORT")
	DATABASE_URL = os.Getenv("DATABASE_URL")
}
