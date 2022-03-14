package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Entry

func LoggerInit() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}
