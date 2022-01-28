package logger

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	devLogger, _ := zap.NewDevelopment()
	logger = devLogger
}

func InfoE(msg string, err error) {
	logger.Info(msg, zap.Error(err))
}

func InfoA(msg string, obj interface{}) {
	logger.Info(msg, zap.Any("data", obj))
}

func Info(msg string) {
	logger.Info(msg)
}
