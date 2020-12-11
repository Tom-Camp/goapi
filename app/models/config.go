package models

import (
	"os"

	"github.com/tom-camp/api.tmc/app/db"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Mongo server ip -> localhost -> 127.0.0.1 -> 0.0.0.0
var server = os.Getenv("DATABASE")

// Database name
var databaseName = os.Getenv("DATABASE_NAME")

func init() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = []string{
		"logs/log.txt",
	}
	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Debug("Sugar start")
}

// Create a connection
var dbConnect = db.NewConnection(server, databaseName)
