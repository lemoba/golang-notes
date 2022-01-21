package config

import (
	"fmt"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDBConnection() *gorm.DB {
	cfg, err := ini.Load(".ini")
	if err != nil {
		panic("Failed to load .ini")
	}

	dbHost := cfg.Section("mysql").Key("DB_HOST")
	dbPort := cfg.Section("mysql").Key("DB_PORT")
	dbUser := cfg.Section("mysql").Key("DB_USER")
	dbPass := cfg.Section("mysql").Key("DB_PASS")
	dbName := cfg.Section("mysql").Key("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database!")
	}
	return db
}

func CloseDBConnection(db *gorm.DB) {
	dbConn, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database!")
	}
	dbConn.Close()
}
