package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"runtime"
)

type logger struct {
	ctx     context.Context
	traceId string
	spanId  string
	pSanId  string
	_logger *zap.Logger
}

func New(ctx context.Context) *logger {
	var traceId, spanId, pSanId string
	if ctx.Value("traceid") != nil {
		traceId = ctx.Value("traceid").(string)
	}
	if ctx.Value("spanid") != nil {
		spanId = ctx.Value("spanid").(string)
	}
	if ctx.Value("pSanId") != nil {
		pSanId = ctx.Value("pSanId").(string)
	}
	return &logger{
		ctx:     ctx,
		traceId: traceId,
		spanId:  spanId,
		pSanId:  pSanId,
		_logger: _logger,
	}
}
func (l *logger) Info(msg string, kv ...interface{}) {
	l.log(zapcore.InfoLevel, msg, kv...)
}
func (l *logger) log(lvl zapcore.Level, msg string, kv ...interface{}) {
	if len(kv)%2 != 0 {
		kv = append(kv, "unknown")
	}
	kv = append(kv, "traceid", l.traceId, "spanid", l.spanId, "pSanid", l.pSanId)
	funcName, file, line := l.getLoggerCallerInfo()
	kv = append(kv, "func", funcName, "file", file, "line", line)
	fields := make([]zap.Field, 0, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		k := fmt.Sprintf("%v", kv[i])
		fields = append(fields, zap.Any(k, kv[i+1]))
	}
	ce := l._logger.Check(lvl, msg)
	ce.Write(fields...)
}
func (l *logger) Debug(msg string, kv ...interface{}) {
	l.log(zapcore.DebugLevel, msg, kv...)
}
func (l *logger) info(msg string, kv ...interface{}) {
	l.log(zapcore.InfoLevel, msg, kv...)
}
func (l *logger) Warn(msg string, kv ...interface{}) {
	l.log(zapcore.WarnLevel, msg, kv...)
}
func (l *logger) Error(msg string, kv ...interface{}) {
	l.log(zapcore.ErrorLevel, msg, kv...)
}
func (l *logger) getLoggerCallerInfo() (funcName, file string, line int) {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return
	}
	file = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return
}
