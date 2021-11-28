package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"proyectos/src/api/config/settings"
	"sync"
)

var (
	once sync.Once
	db   *gorm.DB
)

// Connection returns the connection with database.
func Connection() (*gorm.DB, error) {
	var err error

	var cfg *settings.Data
	cfg = settings.GetData()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", cfg.DBConfig.DBUsername, cfg.DBConfig.DBPassword, cfg.DBConfig.DBHost, cfg.DBConfig.DBName)

	once.Do(func() {
		var sqlDB *sql.DB
		sqlDB, err = sql.Open("mysql", connectionString)
		if err != nil {
			return
		}

		db, err = getInstance(sqlDB)
		if err != nil {
			return
		}
	})

	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, fmt.Errorf("invalid config")
	}
	return db, nil
}

func getInstance(conn *sql.DB) (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
