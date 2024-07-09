package logger

import (
	"log"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	DevLog(args ...interface{})
	Panic(args ...interface{})
	Silly(args ...interface{})
	Fatal(args ...interface{})
}

type LogrusLogger struct {
	*logrus.Logger
}

type StandardLogger struct {
	*log.Logger
}
