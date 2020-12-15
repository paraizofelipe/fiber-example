package setting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Core     CoreSetting
	Database DatabaseSetting
)

type DatabaseSetting struct {
	DBDriver  string
	DBAddress string
}

type CoreSetting struct {
	APIURL string
	Secret string
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panic("Error loading .env file")
	}

	Core = CoreSetting{
		APIURL: os.Getenv("API_URL"),
		Secret: os.Getenv("SECRET"),
	}

	Database = DatabaseSetting{
		DBDriver:  os.Getenv("DB_DRIVER"),
		DBAddress: os.Getenv("DB_ADDRESS"),
	}
}
