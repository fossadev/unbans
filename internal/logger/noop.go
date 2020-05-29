package logger

import (
	"net/http"

	"go.uber.org/zap"
)

type noop struct{}

func Noop() Logger {
	return &noop{}
}

func (l *noop) Debug(msg string, fields ...zap.Field) {
}

func (l *noop) Error(msg string, err error) {
}

func (l *noop) Fatal(msg string, fields ...zap.Field) {
}

func (l *noop) Info(msg string, fields ...zap.Field) {
}

func (l *noop) WithRequest(req *http.Request) Logger {
	return l.With()
}

func (l *noop) With(fields ...zap.Field) Logger {
	return &noop{}
}
