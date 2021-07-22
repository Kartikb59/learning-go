package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger
}

func main() {
	logger := InitLogger()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
	zap.L().Debug("Inside Main")
	var head1 *node = nil
	head1 = insertAtBack(head1, 1)
	insertAtBack(head1, 3)
	insertAtBack(head1, 5)
	insertAtBack(head1, 7)

	var head2 *node = nil
	head2 = insertAtBack(head2, 2)
	insertAtBack(head2, 2)
	insertAtBack(head2, 6)
	insertAtBack(head2, 8)

	var mergedList *node = mergeLL(head1, head2)
	printList(mergedList)
}
