package pkg

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DatabaseHost string
	DatabaseUser string
	DatabasePassword string
	DatabaseName string
	DatabasePort string
}

func LoadConfig() (*Config, error){
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config := &Config{
		ServerPort: getEnv("PORT", "8080"),
		DatabaseHost: getEnv("DATABASE_HOST", "localhost"),
		DatabaseUser: getEnv("DATABASE_USER", "postgres") ,
		DatabasePassword: getEnv("DATABASE_PASSWORD", "") ,
		DatabaseName: getEnv("DATABASE_NAME", "tutorlink_db"),
		DatabasePort: getEnv("DATABASE_PORT", "5432"),
	}

	return config, nil

}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
