package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/yon-module/yon-framework/database"
	"github.com/yon-module/yon-framework/logger"
	"github.com/yon-module/yon-framework/server"
)

const (
	defaultTimeZone = "Asia/Jakarta"
)

func init() {
	// Initial env
	_ = godotenv.Load()

	// Initial logger
	logger.InitLogger()

	setupTimeZone()

	// Initial database
	database.InitDB()

	// Server start
	s := server.NewServer()
	s.Start()
}

func setupTimeZone() {
	timeZone := os.Getenv("yon.server.timezone")
	if timeZone == "" {
		timeZone = defaultTimeZone
	}

	logger.Log.Info().Str("timezone", time.Local.String()).Msg("Loading timezone")
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		panic(err)
	}
	time.Local = loc

}
