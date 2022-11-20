package config

import (
	"log"
	"os"

	"github.com/ezrasitorus77/http-handler/internal/consts"

	"github.com/joho/godotenv"
)

var (
	ServerAddress string
	ServerPort    string
	e             error
)

func init() {
	if e = godotenv.Load(); e != nil {
		log.Fatal(e)
	}

	ServerAddress = os.Getenv("SERVER_ADDRESS")
	if ServerAddress == "" {
		log.Fatal(consts.FataEnvError)
	}

	ServerPort = os.Getenv("SERVER_PORT")
	if ServerPort == "" {
		log.Fatal(consts.FataEnvError)
	}
}
