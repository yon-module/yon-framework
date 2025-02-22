package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339
	Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
	Log = Log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"})
}
