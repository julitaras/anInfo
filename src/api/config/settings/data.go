package settings

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

//instance of config data
var instance *Data

//Data struct
type Data struct {
	Port     string    `json:"port,omitempty"`
	GinMode  string    `json:"gin_mode,omitempty"`
	DBConfig *DBConfig `json:"db_config,omitempty"`
}

//DBConfig struct
type DBConfig struct {
	DBUsername string `json:"db_username,omitempty"`
	DBPassword string `json:"db_password,omitempty"`
	DBHost     string `json:"db_host,omitempty"`
	DBName     string `json:"db_name,omitempty"`
}

//GetData builder
func GetData() *Data {

	if instance == nil {
		if godotenv.Load(".env") != nil {
			log.Fatal("can't load .env")
		}

		port := os.Getenv("PORT")
		env := os.Getenv("GIN_MODE")
		if port == "" {
			port = "8080"
		}

		dbUsername := os.Getenv("DB_USERNAME")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")

		instance = &Data{
			Port:    port,
			GinMode: env,

			DBConfig: &DBConfig{
				DBUsername: dbUsername,
				DBPassword: dbPassword,
				DBHost:     dbHost,
				DBName:     dbName,
			},
		}
	}
	return instance
}
