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
		pc, file, line, ok := runtime.Caller(4)
		if !ok {
			return
		}

		funcName := runtime.FuncForPC(pc).Name()
		shortFunc := funcName[strings.LastIndex(funcName, "/")+1:]

		e.Str("func", shortFunc).
			Str("file", file).
			Int("line", line)
	})

	Log = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger().
		Hook(hook).
		Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02 15:04:05",
		})
}

func DebugLine() zerolog.Logger {
	return Log.Hook()
}
