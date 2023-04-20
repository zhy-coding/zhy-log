// Package options
// @Description:
// @Author:zhy
package options

type LogOption struct {

	// 日志级别
	Level string

	// 输出设置，可选all、console、file 分别为同时打印终端和写入文件、打印终端、写入文件
	OutPut string

	// 输出日志样式，可选json
	Format string

	// 用于使用捕获某个级别日志，将某个级别日志单独写入文件
	FileLevel map[string]LogWriter

	// 写入文件时自定义配置，注意自定义时需要将Used字段设为true
	FileConfig LogWriter
}

// RunLogOption 选择器闭包
type RunLogOption func(r *LogOption) error

/**
  @Description: 初始化选择器
  @param opts 选择
  @return *LogOption
  @return error
*/

func InitLogOption(opts ...RunLogOption) (*LogOption, error) {
	l := &LogOption{
		FileLevel: make(map[string]LogWriter, 0),
	}
	for _, opt := range opts {
		if err := opt(l); err != nil {
			return nil, err
		}
	}
	return l, nil
}

/**
  @Description: 设置日志输出类型
  @param format 可选 json format
  @return RunLogOption
*/

func WithFormat(format string) RunLogOption {
	return func(r *LogOption) error {
		r.Format = format
		return nil
	}
}

/**
  @Description: 设置日志级别
  @param level 可选 debug info warn error panic fatal
  @return RunLogOption
*/

func WithLevel(level string) RunLogOption {
	return func(r *LogOption) error {
		r.Level = level
		return nil
	}
}

/**
  @Description: 设置日志输出位置
  @param output 可选 console all file
  @return RunLogOption
*/

func WithOutPut(output string) RunLogOption {
	// 如果没有设置则默认打印终端
	if output == "" {
		output = "console"
	}
	return func(r *LogOption) error {
		r.OutPut = output
		return nil
	}
}

/**
  @Description: 设置自定义文件配置
  @param fileConfig
  @return RunLogOption
*/

func WithFileConfig(fileConfig LogWriter) RunLogOption {
	return func(r *LogOption) error {
		r.FileConfig = fileConfig
		return nil
	}
}

/**
  @Description: 设置捕获某一级别日志存入文件
  @param l 捕获级别
  @param fileConfig 捕获级别日志的自定义模版，不传则使用默认配置
  @return RunLogOption
*/

func WithFileLevel(l string, fileConfig LogWriter) RunLogOption {
	return func(r *LogOption) error {
		r.FileLevel[l] = fileConfig
		return nil
	}
}
