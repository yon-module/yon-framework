package database

import (
	"time"

	"github.com/yon-module/yon-framework/config"
	"github.com/yon-module/yon-framework/logger"
	"gorm.io/gorm"
	databaseLogger "gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() {
	logger.Log.Info().Msg("Connecting to database...")
	cfg := config.LoadConfig()
	var err error

	// Custom logger untuk GORM
	newLogger := databaseLogger.New(
		&logger.Log, // Gunakan zerolog
		databaseLogger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  databaseLogger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Inisialisasi database
	db, err = gorm.Open(cfg.DSN(), &gorm.Config{Logger: newLogger})
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	logger.Log.Info().Str("db", cfg.DBType).Msg("Database connected successfully")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
