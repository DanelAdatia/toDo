package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB Database
var DB *gorm.DB

var Err error

// DBConfig - Database Configuration
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

// BuildDBConfig builds the database configuration
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		User:     "root",
		Password: "testing321",
		Host:     "127.0.0.1",
		Port:     "3306",
		DBName:   "todoapp",
	}

	return &dbConfig
}

// DBUrl returns the url for the database
func DBUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
