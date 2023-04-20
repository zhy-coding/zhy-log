package logger

import (
	"context"
	"fmt"
	lgr "zhy-log/logrus"
)

const loggerKey = "traceId"

func Debug(args ...interface{}) {
	lgr.GetLogrus().Debug(args...)
}
func Info(args ...interface{}) {
	lgr.GetLogrus().Info(args...)
}
func Warn(args ...interface{}) {
	lgr.GetLogrus().Warn(args...)
}
func Error(args ...interface{}) {
	lgr.GetLogrus().Error(args...)
}

func Fatal(args ...interface{}) {
	lgr.GetLogrus().Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	lgr.GetLogrus().Debug(fmt.Sprintf(format, args...))
}

func Infof(format string, args ...interface{}) {
	lgr.GetLogrus().Info(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...interface{}) {
	lgr.GetLogrus().Warn(fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...interface{}) {
	lgr.GetLogrus().Error(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...interface{}) {
	lgr.GetLogrus().Fatal(fmt.Sprintf(format, args...))
}

type File struct {
	Key   string
	Value interface{}
}

/**
  @Description: 将logger实例存在ctx中
  @param ctx 需要传的ctx
  @param value tranceId
  @return context.Context
*/

func NewContext(ctx *context.Context, file File) {

	c := context.WithValue(*ctx, loggerKey, withValue(*ctx, file))

	*ctx = c

}

/**
  @Description: 返回logger实例
  @param ctx
  @return Log
*/

func WithContext(ctx context.Context) Log {

	value := ctx.Value(loggerKey)

	ctxLogger, ok := value.(Log)

	if !ok {
		return Log{lgr.GetLogrus()}
	}

	return ctxLogger

}

/**
  @Description: 为log添加key-value
  @param ctx
  @param file
  @return Log
*/

func withValue(ctx context.Context, file File) Log {

	log := WithContext(ctx)

	return Log{log.WithField(file.Key, file.Value)}
}
