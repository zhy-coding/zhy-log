package example

import (
	"context"
	"testing"
	"zhy-log/logger"
	"zhy-log/logrus"
	"zhy-log/options"
)

func TestLogrus(t *testing.T) {

	logrus.InitLogrus(options.WithLevel("debug"), options.WithOutPut("console"), options.WithFormat(""))

	c := context.Background()
	logger.NewContext(&c, logger.File{
		Key:   "traceId",
		Value: "123456",
	})

	logger.NewContext(&c, logger.File{
		Key:   "sss",
		Value: "2222",
	})

	logger.WithContext(c).Debug("debug")
	logger.WithContext(c).Info("Info")
	logger.WithContext(c).Error("Error")
	logger.WithContext(c).Warn("Warn")

	logger.WithContext(c).Debugf("deb[%s]ug", "test")

	logger.Debug("test")

	logger.Debugf("te[%v]st", 123)

}
