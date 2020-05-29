package logger

import (
	"fmt"
	"net/http"

	"github.com/fossadev/unbans/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is abstracted, so that we can pass logger.Noop() in tests.
type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Error(msg string, err error)
	// Please don't trigger Fatal's if at all necessary.
	Fatal(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	WithRequest(req *http.Request) Logger
	With(fields ...zap.Field) Logger
}

type logger struct {
	*zap.Logger
}

func New(environment string) (Logger, error) {
	var log *zap.Logger
	var err error

	switch environment {
	case config.EnvDevelopment:
		log, err = zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	default:
		log, err = zap.NewProduction()
	}
	if err != nil {
		return nil, err
	}

	return &logger{Logger: log}, nil
}

func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, fields...)
}

func (l *logger) Error(msg string, err error) {
	l.Logger.Error(msg, zap.Error(err))
}

func (l *logger) Fatal(msg string, fields ...zap.Field) {
	l.Logger.Fatal(msg, fields...)
}

func (l *logger) Info(msg string, fields ...zap.Field) {
	l.Logger.Info(msg, fields...)
}

func (l *logger) WithRequest(req *http.Request) Logger {
	return l.With(zap.String("request-uri", fmt.Sprintf("%s %s", req.Method, req.RequestURI)))
}

func (l *logger) With(fields ...zap.Field) Logger {
	return &logger{Logger: l.Logger.With(fields...)}
}
