package logger

import "fmt"

// will add more logic for formatting and stuff
func (l StandardLogger) Info(args ...interface{}) {
	l.Logger.Println(args...)
}

func (l StandardLogger) Debug(args ...interface{}) {
	l.Logger.Println(args...)
}

func (l StandardLogger) Error(args ...interface{}) {
	l.Logger.Println(args...)
}

func (l StandardLogger) DevLog(args ...interface{}) {
	l.Logger.Println(args...)
}

func (l StandardLogger) Panic(args ...interface{}) {
	l.Logger.Panic(args...)
}

func (l StandardLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}

func (l StandardLogger) Silly(args ...interface{}) {
	fmt.Println(args...)
}
