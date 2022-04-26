package utils

import (
	"errors"
	"github.com/CatMoe/Ayaka/config"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"os"
)

var logrusLogger *logrus.Entry

type LogLevel uint32

const (
	Panic LogLevel = iota
	Fatal
	Error
	Warn
	Info
	Debug
)

func logUtilInit() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	sentry.Init(sentry.ClientOptions{
		Dsn:              config.SENTRY_DSN,
		Debug:            true,
		AttachStacktrace: true,
	})
}

func Log(level LogLevel, message string) {
	switch level {
	case Panic:
		logrus.Panic(message)
		sentry.CaptureException(errors.New(message))
	case Fatal:
		logrus.Fatal(message)
		sentry.CaptureException(errors.New(message))
	case Error:
		logrus.Error(message)
		sentry.CaptureException(errors.New(message))
	case Warn:
		logrus.Warn(message)
		sentry.CaptureMessage(message)
	case Info:
		logrus.Info(message)
		sentry.CaptureMessage(message)
	case Debug:
		logrus.Debug(message)
		sentry.CaptureMessage(message)
	}
}
