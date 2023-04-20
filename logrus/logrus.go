// Package logrus
// @Description:
// @Author:zhy
package logrus

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"zhy-log/options"
)

var log = logrus.New()

/**
  @Description: 初始化logrus配置
  @param outType 日志输出类型，分为json和正常
  @param level 日志级别
  @param types 日志输出模式
*/

func InitLogrus(opts ...options.RunLogOption) {

	logOption, _ := options.InitLogOption(opts...)

	parseLevel, _ := logrus.ParseLevel(logOption.Level)

	log.SetFormatter(&LogFormatter{})

	if logOption.Format == "json" {

		log.SetFormatter(&logrus.JSONFormatter{})
	}

	log.SetLevel(parseLevel)

	log.SetReportCaller(true)

	SetOut(log, logOption)

}

/**
  @Description: 设置输出类型
  @param logger
  @param types 输出类型
  @param logFile 日志存放路径
*/

func SetOut(logger *logrus.Logger, l *options.LogOption) {

	mw := io.MultiWriter()

	switch l.OutPut {
	case "all":
		mw = io.MultiWriter(os.Stdout, options.GetWriter(l.FileConfig))
	case "console":
		mw = io.MultiWriter(os.Stdout)
	case "file":
		mw = io.MultiWriter(options.GetWriter(l.FileConfig))
	}

	logger.SetOutput(mw)
}

/**
  @Description: 获取包内log
  @return *logrus.Logger
*/

func GetLogrus() *logrus.Entry {
	return log.WithContext(context.Background())
}
