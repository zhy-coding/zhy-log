// Package logrus
// @Description: 自定义日志模版
// @Author:zhy
package logrus

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

//  @Description: 自定义一个日志格式化器
//  @author zhy
//  @date 2023-04-04 10:56:37

type LogFormatter struct {
}

//  @Description: 实现 Formatter 接口的 Format 方法
//  @param entry
//  @return []byte
//  @return error

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	var msg string

	pc, _, line, _ := runtime.Caller(6)

	// 获取调用函数的对象
	fn := runtime.FuncForPC(pc)

	// 获取调用函数的名称
	name := fn.Name()

	level := unifyLevel(entry.Level.String())

	if entry.HasCaller() {

		msg = fmt.Sprintf("[%s] [%s] [%s:%d] %s\n",
			timestamp, level, name, line, entry.Message)

		if len(entry.Data) > 0 {
			var b *bytes.Buffer

			if entry.Buffer != nil {
				b = entry.Buffer
			} else {
				b = &bytes.Buffer{}
			}

			f.appendNormal(b, timestamp)

			f.appendNormal(b, level)

			for key, value := range entry.Data {
				f.appendKeyValue(b, key, value)
			}

			f.appendNormal(b, fmt.Sprintf("%s:%d", name, line))

			b.WriteString(" ")

			b.WriteString(entry.Message)

			b.WriteString(" \n")

			msg = b.String()
			//
			//msg = fmt.Sprintf("[%s] [%s] [%s] [%s:%d] %s\n",
			//	timestamp, level, entry.Data, name, line, entry.Message)
		}

	} else {
		msg = fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, entry.Message)
	}

	b.WriteString(msg)

	return b.Bytes(), nil
}

func (f *LogFormatter) appendNormal(b *bytes.Buffer, normal string) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString("[")
	b.WriteString(normal)
	b.WriteString("]")

}

func (f *LogFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString("[")

	b.WriteString(key)
	b.WriteByte('=')
	//f.appendValue(b, value)
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}
	b.WriteString(stringVal)
	b.WriteString("]")

}

//  @Description: 统一日志级别格式
//  @param level
//  @return string

func unifyLevel(level string) string {
	switch level {
	case "debug":
		return colorize("DEBG", "green")
	case "info":
		return colorize("INFO", "blue")
	case "error":
		return colorize("EROR", "red")
	case "warning":
		return colorize("WARN", "yellow")
	case "fatal":
		return colorize("FATL", "red")
	case "panic":
		return colorize("PNIC", "red")
	}
	return ""
}

func colorize(s string, color string) string {
	colors := map[string]string{
		"reset":  "\033[0m",
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"gray":   "\033[37m",
	}
	return fmt.Sprintf("%s%s%s", colors[color], s, colors["reset"])
}
