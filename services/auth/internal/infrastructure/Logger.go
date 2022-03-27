package infrastructure

import (
	"go.uber.org/zap"
)

type Logger struct {
	log *zap.SugaredLogger
}

func NewLogger() *Logger {
	log, _ := zap.NewProduction()
	sugarLog := log.Sugar()

	return &Logger{log: sugarLog}
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.log.Infof(template, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.log.Warnf(template, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.log.Errorf(template, args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.log.Fatalf(template, args...)
}

func (l *Logger) Infow(template string, keyValuePairs ...interface{}) {
	l.log.Infow(template, keyValuePairs...)
}

func (l *Logger) Warnw(template string, keyValuePairs ...interface{}) {
	l.log.Warnw(template, keyValuePairs...)
}

func (l *Logger) Errorw(template string, keyValuePairs ...interface{}) {
	l.log.Errorw(template, keyValuePairs...)
}

func (l *Logger) Fatalw(template string, keyValuePairs ...interface{}) {
	l.log.Fatalw(template, keyValuePairs...)
}
