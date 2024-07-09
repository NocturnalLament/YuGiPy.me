package logger

func (l LogrusLogger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l LogrusLogger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

func (l LogrusLogger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}

func (l LogrusLogger) DevLog(args ...interface{}) {
	l.Logger.Debug(args...)
}

func (l LogrusLogger) Panic(args ...interface{}) {
	l.Logger.Panic(args...)
}

func (l LogrusLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}

func (l LogrusLogger) Silly(args ...interface{}) {
	l.Logger.Debug(args...)
}
