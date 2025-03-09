package database

import (
	"time"

	"github.com/yon-module/yon-framework/logger"
	"gorm.io/gorm"
	databaseLogger "gorm.io/gorm/logger"
)

var db *gorm.DB
var tableMigration []interface{}

func MigrationRegister(tables ...interface{}) {
	tableMigration = append(tableMigration, tables...)
}

func InitDB() {
	logger.Log.Info().Msg("🏗️ Connecting to database...")
	cfg := LoadConfig()

	logger.Log.Info().
		Str("db", cfg.DBType).
		Str("dbHost", cfg.DBHost).
		Str("dbPort", cfg.DBPort).
		Str("dbName", cfg.DBName).
		Str("dbUser", cfg.DBUser).
		Str("dbPassword", "xxxx").Msg("Trying...")

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
		logger.Log.Fatal().Err(err).Msg("❌ Failed to connect to database")
	}

	logger.Log.Info().Msg("✅ Database connected successfully")
	logger.Log.Info().Msg("🚀 Running migration...")
	_ = GetDB().AutoMigrate(tableMigration...)
	logger.Log.Info().Msg("✅ Migration completed successfully!")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
