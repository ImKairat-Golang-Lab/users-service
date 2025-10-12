package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	log *zap.Logger
}

func NewZapLogger() *ZapLogger {
	log, err := zap.NewProduction()
	if err != nil {
		// (!) Что-бы на старте уловить ощибку связанное с настройкой логгера
		panic(err)
	}

	return &ZapLogger{
		log: log,
	}
}

func mapToFields(fields map[string]any) []zap.Field {
	var zapFields []zap.Field
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}

	return zapFields
}

func (zl *ZapLogger) Debug(msg string, fields map[string]any) {
	zl.log.Debug(msg, mapToFields(fields)...)
}

func (zl *ZapLogger) Info(msg string, fields map[string]any) {
	zl.log.Info(msg, mapToFields(fields)...)
}

func (zl *ZapLogger) Warn(msg string, fields map[string]any) {
	zl.log.Warn(msg, mapToFields(fields)...)
}

func (zl *ZapLogger) Error(msg string, fields map[string]any) {
	zl.log.Error(msg, mapToFields(fields)...)
}
