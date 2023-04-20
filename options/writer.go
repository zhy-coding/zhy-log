// Package options
// @Description: writer
// @Author:zhy
package options

import (
	"github.com/natefinch/lumberjack"
	"io"
)

//  @Description: 自定义写入文件的日志存储
//  @author zhy
//  @date 2023-04-10 15:51:38

type LogWriter struct {
	// 文件名
	FileName string

	// 单文件最大数
	MaxSize int

	// 最大文件保留数
	MaxBackups int

	// 文件保存天数
	MaxAge int

	// 是否压缩
	Compress bool

	// 是否使用自定义模版，false使用默认Writer
	Used bool
}

// 默认writer不自定义时使用默认writer
var defaultWriter = &lumberjack.Logger{
	Filename:   "./logs.txt",
	MaxSize:    10,    //最大M数，超过则切割
	MaxBackups: 5,     //最大文件保留数，超过就删除最老的日志文件
	MaxAge:     30,    //保存30天
	Compress:   false, //是否压缩
}

/**
  @Description: 获取writer
  @param w 自定义配置
  @return io.Writer
*/

func GetWriter(w LogWriter) io.Writer {
	if !w.Used {
		return defaultWriter
	}
	return &lumberjack.Logger{
		Filename:   w.FileName,
		MaxSize:    w.MaxSize,    //最大M数，超过则切割
		MaxBackups: w.MaxBackups, //最大文件保留数，超过就删除最老的日志文件
		MaxAge:     w.MaxAge,     //保存30天
		Compress:   w.Compress,   //是否压缩
	}
}
