package server

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/yon-module/yon-framework/logger"
	"os"
)

func initialLoggerSentry(app *gin.Engine) {
	if os.Getenv("yon.logger.sentry") == "true" {
		dsn := os.Getenv("yon.logger.sentry.dsn")
		options := sentry.ClientOptions{
			Dsn:              dsn,
			AttachStacktrace: true,
			SampleRate:       1.0,
			EnableTracing:    true,
			TracesSampleRate: 1.0,
			ServerName:       "academic-api@v1.0.0",
			Release:          "academic-api@v1.0.0",
			Environment:      "production",
		}
		if err := sentry.Init(options); err != nil {
			logger.Log.Warn().Msgf("Sentry initialization failed: %v", err)
		}

		app.Use(sentrygin.New(sentrygin.Options{
			Repanic: true,
		}))
	}
}
