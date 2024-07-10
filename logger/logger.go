package logger

import (
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type Level int

const (
	Info Level = iota
	Debug
	Error
	DevLog
	Panic
	Silly
	Fatal
)

func (l Level) String() string {
	switch l {
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	case Error:
		return "ERROR"
	case DevLog:
		return "DEVLOG"
	case Panic:
		return "PANIC"
	case Silly:
		return "SILLY"
	case Fatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

type Logger interface {
	Info(message string)
	Debug(message string)
	Error(message string)
	DevLog(message string)
	Panic(message string)
	Silly(message string)
	Fatal(message string)
}

func Colorize(level Level) (func(a ...interface{}) string, func(a ...interface{}) string) {
	var colorFunc func(a ...interface{}) string
	var levelBase func(a ...interface{}) string
	switch level {
	case Info:
		colorFunc = color.New(color.FgGreen).SprintFunc()
		levelBase = color.New(color.FgGreen).Add(color.Bold).Add(color.Underline).SprintFunc()
	case Debug:
		colorFunc = color.New(color.FgBlue).SprintFunc()
		levelBase = color.New(color.FgBlue).Add(color.Bold).Add(color.Underline).SprintFunc()
	case Error:
		colorFunc = color.New(color.FgRed).SprintFunc()
		levelBase = color.New(color.FgRed).Add(color.Bold).Add(color.Underline).SprintFunc()
	case DevLog:
		colorFunc = color.New(color.FgYellow).SprintFunc()
		levelBase = color.New(color.FgYellow).Add(color.Bold).Add(color.Underline).SprintFunc()
	case Panic:
		colorFunc = color.New(color.FgMagenta).SprintFunc()
		levelBase = color.New(color.FgMagenta).Add(color.Bold).Add(color.Underline).SprintFunc()
	}
	return colorFunc, levelBase
}
func LoggerFormatter(level Level, message string) string {
	time := time.Now().Format("2006-01-02 15:04:05")
	color, colorBase := Colorize(level)
	outputBase := level.String()
	output := colorBase(outputBase) + " " + color(time+" - "+message) + "\n"
	return output
}

func WriteLogToFile(path string, level Level, message string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	time := time.Now().Format("2006-01-02 15:04:05")
	outputBase := level.String() + " " + time + " - " + message + "\n"
	_, err = file.WriteString(outputBase)
	if err != nil {
		log.Fatal(err)
	}
}

type LogrusLogger struct {
	*logrus.Logger
	Level
	LogFilePath string
}

type StandardLogger struct {
	*log.Logger
	Level
	LogFilePath string
}
