package config

import (
	"fmt"
	"os"

	"github.com/yon-module/yon-framework/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ConfigDatabase interface {
	DSN() gorm.Dialector
}

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	DBType     string
}

func LoadConfig() Config {
	return Config{
		DBUser:     os.Getenv("yon.database.user"),
		DBPassword: os.Getenv("yon.database.password"),
		DBName:     os.Getenv("yon.database.db"),
		DBHost:     os.Getenv("yon.database.host"),
		DBPort:     os.Getenv("yon.database.port"),
		DBType:     os.Getenv("yon.database.type"),
	}
}

func (c Config) DSN() gorm.Dialector {
	var dsn string
	var dialector gorm.Dialector

	switch c.DBType {
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
		dialector = postgres.Open(dsn)

	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
		dialector = mysql.Open(dsn)

	case "sqlite":
		dsn = c.DBName
		dialector = sqlite.Open(dsn)

	default:
		logger.Log.Fatal().Msg("Unsupported database type")
	}

	return dialector
}
