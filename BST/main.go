package main

import (
	pack "BST/btree"

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
	var root *pack.Node = nil
	var myRoot *pack.Node = pack.InsertNode(root, 3)
	pack.InsertNode(myRoot, 2)
	pack.InsertNode(myRoot, 77)
	pack.InsertNode(myRoot, 45)
	pack.InsertNode(myRoot, 20)
	pack.InsertNode(myRoot, 712)
	pack.InsertNode(myRoot, 172)
	pack.InsertNode(myRoot, 572)
	pack.InsertNode(myRoot, 1)

	pack.Inorder(myRoot)

}
