package server

import (
	"github.com/smallnest/rpcx/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func InitLog() *zap.Logger {
	logger, err := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalColorLevelEncoder, // INFO

			TimeKey:    "time",
			EncodeTime: zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),

			CallerKey:        "caller",
			EncodeCaller:     zapcore.ShortCallerEncoder,
			ConsoleSeparator: " ",
			FunctionKey:      "",
		},
	}.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

func TestCustomLogger(t *testing.T) {

	zapLogger := &log.ZapLogger{}
	zapLogger.Set(InitLog())
	log.SetLogger(zapLogger)
	log.Info("test 2")
	log.Infof("%d", 123)

	log.Warnf("%d", 123)
	log.Warn("num:", 123)

	log.Debugf("%d", 123)
	log.Debug("num:", 123)

	log.Errorf("%d", 123)
	log.Error("num:", 123)

	log.Panicf("%d", 123)
	log.Panic("num:", 123)

	log.Fatalf("%d", 123)
	log.Fatal("num:", 123)
}
