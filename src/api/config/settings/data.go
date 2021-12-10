package settings

import (
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
	DBPort     string `json:"db_port,omitempty"`
}

//GetData builder
func GetData() *Data {

	if instance == nil {
		port := os.Getenv("PORT")
		env := os.Getenv("GIN_MODE")
		if port == "" {
			port = "8080"
		}

		dbUsername := "lzdsmcqrqitpns"
		dbPassword := "759a9d75358226257eaf1226251e085e9a22cb4920cfa145cbc6904c460ba68f"
		dbHost := "ec2-52-71-217-158.compute-1.amazonaws.com"
		dbName := "d36jfrm60cfdgd"
		dbPort := "5432"

		instance = &Data{
			Port:    port,
			GinMode: env,

			DBConfig: &DBConfig{
				DBUsername: dbUsername,
				DBPassword: dbPassword,
				DBHost:     dbHost,
				DBName:     dbName,
				DBPort:     dbPort,
			},
		}
	}
	return instance
}
