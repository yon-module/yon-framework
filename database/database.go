package database

import (
	"fmt"
	"time"

	"github.com/yon-module/yon-framework/config"
	"github.com/yon-module/yon-framework/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	databaseLogger "gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() {
	logger.Log.Info().Msg("Connecting to database...")
	cfg := config.LoadConfig()
	var err error
	var dsn string
	var dialector gorm.Dialector

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

	// Pilih database driver berdasarkan DB_TYPE
	switch cfg.DBType {
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
		dialector = postgres.Open(dsn)

	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
		dialector = mysql.Open(dsn)

	case "sqlite":
		dsn = cfg.DBName // Nama file SQLite
		dialector = sqlite.Open(dsn)

	default:
		logger.Log.Fatal().Msg("Unsupported database type")
	}

	// Inisialisasi database
	db, err = gorm.Open(dialector, &gorm.Config{Logger: newLogger})
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	logger.Log.Info().Str("db", cfg.DBType).Msg("Database connected successfully")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
