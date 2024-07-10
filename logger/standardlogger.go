package logger

// will add more logic for formatting and stuff
func (l StandardLogger) Info(message string) {
	level := Info
	WriteLogToFile(l.LogFilePath, level, message)
	l.Logger.Println(LoggerFormatter(level, message))
}

func (l StandardLogger) Debug(message string) {
	l.Logger.Println(message)
}

func (l StandardLogger) Error(message string) {
	l.Logger.Println(message)
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
	l.Logger.Println(args...)
}
