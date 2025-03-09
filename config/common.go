package config

import (
	"github.com/joho/godotenv"
	"github.com/yon-module/yon-framework/database"
	"github.com/yon-module/yon-framework/logger"
	"github.com/yon-module/yon-framework/server"
)

func init() {
	// Initial env
	_ = godotenv.Load()

	// Initial logger
	logger.InitLogger()

	// Initial database
	database.InitDB()

	// Server start
	s := server.NewServer()
	s.Start()
}
