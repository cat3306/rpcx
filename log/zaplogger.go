package log

import "go.uber.org/zap"

type ZapLogger struct {
	logger *zap.Logger
}

func (z *ZapLogger) Set(l *zap.Logger) {
	z.logger = l
}
func (z *ZapLogger) Debug(v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Debug(v...)
}
func (z *ZapLogger) Debugf(format string, v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Debugf(format, v...)
}
func (z *ZapLogger) Info(v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Info(v...)

}
func (z *ZapLogger) Infof(format string, v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Infof(format, v...)
}
func (z *ZapLogger) Warnf(format string, v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Warnf(format, v...)
}
func (z *ZapLogger) Warn(v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Warn(v...)
}
func (z *ZapLogger) Errorf(format string, v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Errorf(format, v...)
}
func (z *ZapLogger) Error(v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Error(v...)
}
func (z *ZapLogger) Panic(v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Panic(v...)
}
func (z *ZapLogger) Panicf(format string, v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Panicf(format, v...)
}
func (z *ZapLogger) Fatalf(format string, v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Fatalf(format, v...)
}
func (z *ZapLogger) Fatal(v ...interface{}) {
	z.logger.Sugar().WithOptions(zap.AddCallerSkip(2)).Fatal(v...)
}
