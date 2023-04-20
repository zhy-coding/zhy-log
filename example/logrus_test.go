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

	c := logger.NewContext(context.Background(), "1234567")

	logger.WithContext(c).Debug("debug")
	logger.WithContext(c).Info("Info")
	logger.WithContext(c).Error("Error")
	logger.WithContext(c).Warn("Warn")

	logger.WithContext(c).Debugf("deb[%s]ug", "test")

	logger.Debug("test")

	logger.Debugf("te[%v]st", 123)

}
