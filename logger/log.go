package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

//  @Description: 该结构体存在是为了确保获取行数一致所多加的一层
//  @author zhy
//  @date 2023-04-12 14:22:59

type Log struct {
	lgr *logrus.Entry
}

func (l Log) Debug(args ...interface{}) {
	l.lgr.Debug(args...)
}

func (l Log) Info(args ...interface{}) {
	l.lgr.Info(args...)
}

func (l Log) Warn(args ...interface{}) {
	l.lgr.Warn(args...)
}

func (l Log) Error(args ...interface{}) {
	l.lgr.Error(args...)
}

func (l Log) Fatal(args ...interface{}) {
	l.lgr.Fatal(args...)
}

func (l Log) Debugf(format string, args ...interface{}) {
	l.lgr.Debug(fmt.Sprintf(format, args...))
}

func (l Log) Infof(format string, args ...interface{}) {
	l.lgr.Info(fmt.Sprintf(format, args...))
}

func (l Log) Warnf(format string, args ...interface{}) {
	l.lgr.Warn(fmt.Sprintf(format, args...))
}

func (l Log) Errorf(format string, args ...interface{}) {
	l.lgr.Error(fmt.Sprintf(format, args...))
}

func (l Log) Fatalf(format string, args ...interface{}) {
	l.lgr.Fatal(fmt.Sprintf(format, args...))
}
