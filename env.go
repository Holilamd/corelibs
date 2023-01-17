package corelibs

import (
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(key string) string {

	// load .env file
	_ = godotenv.Load(".env")

	// if err != nil {
	// 	CommonLogger().Info("can't load .env file, get os key")
	// }

	return os.Getenv(key)
}
