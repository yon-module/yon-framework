package database

import (
	"github.com/yon-module/yon-framework/config"
	"github.com/yon-module/yon-framework/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.LoadConfig()
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal().Str("Failed to connect to database: ", err.Error())
	}
	logger.Log.Info().Msg("Database connected")
}
