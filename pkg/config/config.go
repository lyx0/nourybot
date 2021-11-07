package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Username  string
	Oauth     string
	BotUserId string
	MongoURI  string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	cfg := &Config{
		Username:  os.Getenv("TWITCH_USER"),
		Oauth:     os.Getenv("TWITCH_PASS"),
		BotUserId: os.Getenv("BOT_USER_ID"),
		MongoURI:  os.Getenv("MONGO_URI"),
	}

	log.Info("Config loaded succesfully")

	return cfg
}

// Only for tests
func LoadConfigTest() {
	os.Setenv("TEST_VALUE", "xDLUL420")
	// defer os.Unsetenv("TEST_VALUE")
}
