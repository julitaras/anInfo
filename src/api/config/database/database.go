package database

import (
	"database/sql"
	"fmt"
	"log"
	"proyectos/src/api/config/settings"
	"sync"
)

var (
	once sync.Once
	db   *sql.DB
)

//Connection returns a connection to DB
func Connection() (*sql.DB, error) {
	var err error
	log.Println("*Database* Connection start *new instance*")
	var cfg *settings.Data
	cfg = settings.GetData()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", cfg.DBConfig.DBUsername, cfg.DBConfig.DBPassword, cfg.DBConfig.DBHost, cfg.DBConfig.DBName)

	once.Do(func() {
		db, err = sql.Open("mysql", connectionString)
	})

	if err != nil {
		log.Fatalf("*ERROR* *Database* "+err.Error(), err)
		return nil, err
	}

	if errPing := db.Ping(); errPing != nil && cfg.GinMode != "test" {
		err = errPing
		log.Fatalf("*ERROR* *Database* Error pinging database: "+errPing.Error(), err)
		return nil, err
	}

	return db, nil
}
