package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

//  @Description: 该结构体存在是为了确保获取行数一致所多加的一层
//  @author zhy
//  @date 2023-04-12 14:22:59

type Log struct {
	*logrus.Entry
}

func (l Log) Debug(args ...interface{}) {
	l.Entry.Debug(args...)
}

func (l Log) Info(args ...interface{}) {
	l.Entry.Info(args...)
}

func (l Log) Warn(args ...interface{}) {
	l.Entry.Warn(args...)
}

func (l Log) Error(args ...interface{}) {
	l.Entry.Error(args...)
}

func (l Log) Fatal(args ...interface{}) {
	l.Entry.Fatal(args...)
}

func (l Log) Debugf(format string, args ...interface{}) {
	l.Entry.Debug(fmt.Sprintf(format, args...))
}

func (l Log) Infof(format string, args ...interface{}) {
	l.Entry.Info(fmt.Sprintf(format, args...))
}

func (l Log) Warnf(format string, args ...interface{}) {
	l.Entry.Warn(fmt.Sprintf(format, args...))
}

func (l Log) Errorf(format string, args ...interface{}) {
	l.Entry.Error(fmt.Sprintf(format, args...))
}

func (l Log) Fatalf(format string, args ...interface{}) {
	l.Entry.Fatal(fmt.Sprintf(format, args...))
}
