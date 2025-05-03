package logger

import (
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339
	hook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		pc, file, line, ok := runtime.Caller(3)
		if !ok {
			return
		}

		funcName := runtime.FuncForPC(pc).Name()
		shortFunc := funcName[strings.LastIndex(funcName, "/")+1:]

		e.Str("Func", shortFunc).
			Str("File", file).
			Int("On Line", line)
	})

	Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
	Log = Log.Hook(hook).Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
	})

	Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
	Log = Log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"})
}
